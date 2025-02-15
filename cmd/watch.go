package cmd

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/yairp7/gotcher/internal/events"
	"github.com/yairp7/gotcher/internal/utils"
)

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.Flags().StringSlice("events", nil, "The events we want to be triggered by: WRITE, REMOVE, RENAME and CHMOD (comma separated)")
	watchCmd.Flags().String("pattern", "", "The pattern of the files we want to be triggerd by their events")
	watchCmd.Flags().String("cmd", "", "The command to run when the event occours")
}

func events2Ops(eventsFlags []string) ([]fsnotify.Op, error) {
	ops := make([]fsnotify.Op, 0)
	for _, event := range eventsFlags {
		op, err := utils.Name2Op(event)
		if err != nil {
			return nil, err
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func getDirsToWatch(path string) []string {
	subDirs, err := utils.ListDirs(path)
	if err != nil {
		ExitWithError(fmt.Errorf("failed creating watcher - %v", err))
	}
	dirs := make([]string, len(subDirs)+1)
	dirs[0] = path
	copy(dirs[1:], subDirs)
	return dirs
}

func runWatcher(ctx context.Context, dirs []string) (<-chan fsnotify.Event, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	eventsChan := make(chan fsnotify.Event)
	go func(eventsChan chan fsnotify.Event) {
		defer watcher.Close()
		defer close(eventsChan)

		for _, dir := range dirs {
			path := fmt.Sprintf("./%s", dir)
			err = watcher.Add(path)
			if err != nil {
				ExitWithError(fmt.Errorf("failed creating watcher - %v", err))
			}
			log.Printf("Added Watcher for %s", path)
		}

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				eventsChan <- event
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("error: %v", err)
			case <-ctx.Done():
				return
			}
		}
	}(eventsChan)

	return eventsChan, nil
}

func onResult(result events.Result) {
	log.Println(color.BlueString(fmt.Sprintf("Modified: %s", result.Path)))

	if result.Err != nil {
		log.Println(color.RedString(fmt.Sprintf("%v", result.Err)))
	}
}

var watchCmd = &cobra.Command{
	Use:   "watch <path> <action>",
	Short: "Start watching files and add actions to the events",
	Long:  `Start watching files and add actions to the events`,
	Run: func(cmd *cobra.Command, args []string) {
		var ops []fsnotify.Op
		var pattern string
		var execCmd string

		eventsFlags, err := cmd.Flags().GetStringSlice("events")
		if err != nil {
			ExitWithError(fmt.Errorf("must provide events types - %v", err))
		}

		if eventsFlags != nil {
			ops, err = events2Ops(eventsFlags)
			if err != nil {
				ExitWithError(fmt.Errorf("failed parsing events types - %v", err))
			}
		}

		pattern, err = cmd.Flags().GetString("pattern")
		if err != nil {
			ExitWithError(fmt.Errorf("bad file pattern"))
		}

		execCmd, err = cmd.Flags().GetString("cmd")
		if err != nil {
			ExitWithError(fmt.Errorf("bad command"))
		}

		if len(args) < 1 {
			ExitWithError(fmt.Errorf("wrong usage of the command, try something like this: %s", cmd.Use))
		}

		path := args[0]
		if len(path) == 10 {
			ExitWithError(fmt.Errorf("must provide a path"))
		} else if ok, err := utils.Exists(path); !ok || err != nil {
			ExitWithError(fmt.Errorf("must provide a valid path to watch"))
		}

		dirs := getDirsToWatch(path)

		ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
		defer cancelFunc()

		eventsProcessor, err := events.NewEventProcessor(ops, pattern, execCmd)
		if err != nil {
			ExitWithError(fmt.Errorf("failed creating processor - %v", err))
		}
		defer eventsProcessor.Close()

		eventsChan, err := runWatcher(ctx, dirs)
		if err != nil {
			ExitWithError(fmt.Errorf("failed creating watcher - %v", err))
		}

		resultsChan := eventsProcessor.Run(ctx, eventsChan)

		for {
			select {
			case result := <-resultsChan:
				onResult(result)
			case <-ctx.Done():
				log.Print("\r")
				log.Println("Shutting down...")
				return
			}
		}
	},
}
