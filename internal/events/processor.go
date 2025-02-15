package events

import (
	"context"
	"regexp"
	"slices"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/yairp7/gotcher/internal/utils"
)

type Result struct {
	Path string
	Err  error
}

type EventProcessor struct {
	resultsChan  chan Result
	ops          []fsnotify.Op
	fileRegexp   *regexp.Regexp
	commandToRun string
}

func NewEventProcessor(ops []fsnotify.Op, pattern string, commandToRun string) (*EventProcessor, error) {
	r, err := regexp.Compile(pattern)
	return &EventProcessor{
		resultsChan:  nil,
		ops:          ops,
		fileRegexp:   r,
		commandToRun: commandToRun,
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

			if len(ep.commandToRun) == 0 {
				continue
			}

			result := Result{Path: event.Name}
			commandToRun := ep.replaceArgumentsIfNeeded(ep.commandToRun, event)
			if err := utils.ExecShell(ctx, commandToRun); err != nil {
				break
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
