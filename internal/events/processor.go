package events

import (
	"context"
	"regexp"
	"slices"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/yairp7/gotcher/internal/utils"
)

type Handler interface {
	Handle(ctx context.Context, event fsnotify.Event) error
}

type Result struct {
	Path string
	Err  error
}

type EventProcessor struct {
	resultsChan  chan Result
	ops          []fsnotify.Op
	fileRegexp   *regexp.Regexp
	commandToRun string
	handlers     []Handler
	logger       utils.Logger
}

func NewEventProcessor(ops []fsnotify.Op, pattern string, commandToRun string, logger utils.Logger, handlers ...Handler) (*EventProcessor, error) {
	r, err := regexp.Compile(pattern)
	return &EventProcessor{
		resultsChan:  nil,
		ops:          ops,
		fileRegexp:   r,
		commandToRun: commandToRun,
		handlers:     handlers,
		logger:       logger,
	}, err
}

func (ep *EventProcessor) replaceArgumentsIfNeeded(commandToRun string, event fsnotify.Event) string {
	commandToRun = strings.ReplaceAll(commandToRun, "#[file]", event.Name)
	commandToRun = strings.ReplaceAll(commandToRun, "#[op]", utils.Op2Name(event.Op))
	return commandToRun
}

func (ep *EventProcessor) run(ctx context.Context, eventsChan <-chan fsnotify.Event) {
	for {
		select {
		case event := <-eventsChan:
			if len(ep.ops) > 0 && !slices.Contains(ep.ops, event.Op) {
				continue
			}

			if ep.fileRegexp != nil && !ep.fileRegexp.MatchString(event.Name) {
				continue
			}

			for _, handler := range ep.handlers {
				if err := handler.Handle(ctx, event); err != nil {
					ep.resultsChan <- Result{Path: event.Name, Err: err}
				}
			}

			result := Result{Path: event.Name}

			if len(ep.commandToRun) == 0 {
				ep.resultsChan <- result
				continue
			}

			commandToRun := ep.replaceArgumentsIfNeeded(ep.commandToRun, event)
			ep.logger.Debug("Executing command: %s", commandToRun)
			if err := utils.ExecShell(ctx, commandToRun); err != nil {
				ep.logger.Error("Failed to execute command: %v", err)
				ep.resultsChan <- Result{Path: event.Name, Err: err}
				continue
			}
			ep.resultsChan <- result
		case <-ctx.Done():
			return
		}
	}
}

func (ep *EventProcessor) Run(ctx context.Context, eventsChan <-chan fsnotify.Event) <-chan Result {
	ep.resultsChan = make(chan Result)
	go ep.run(ctx, eventsChan)
	return ep.resultsChan
}

func (ep *EventProcessor) Close() {
	close(ep.resultsChan)
}
