package utils

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func WatchDir(watcher *fsnotify.Watcher, dir string) error {
	err := watcher.Add(dir)
	if err != nil {
		return fmt.Errorf("failed adding watcher for %s - %v", dir, err)
	}
	return nil
}

func UnwatchDir(watcher *fsnotify.Watcher, dir string) error {
	err := watcher.Remove(dir)
	if err != nil && err != fsnotify.ErrNonExistentWatch {
		return fmt.Errorf("failed removing watcher for %s - %v", dir, err)
	}
	return nil
}
