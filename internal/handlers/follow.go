package handlers

import (
	"context"
	"os"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/yairp7/gotcher/internal/utils"
)

type FollowHandler struct {
	watcher *fsnotify.Watcher
	logger  utils.Logger
}

func NewFollowHandler(watcher *fsnotify.Watcher, logger utils.Logger) *FollowHandler {
	return &FollowHandler{watcher: watcher, logger: logger}
}

func (h *FollowHandler) isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func (h *FollowHandler) isExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (h *FollowHandler) Handle(ctx context.Context, event fsnotify.Event) error {
	h.logger.Debug("Event: %v", event)
	if event.Op == fsnotify.Create && h.isDir(event.Name) {
		if err := utils.WatchDir(h.watcher, event.Name); err != nil {
			return err
		}
		h.logger.Info("Started watching %s", color.GreenString(event.Name))
	} else if (event.Op == fsnotify.Remove || event.Op == fsnotify.Rename) && !h.isExists(event.Name) {
		if err := utils.UnwatchDir(h.watcher, event.Name); err != nil {
			return err
		}
		h.logger.Info("Stopped watching %s", color.GreenString(event.Name))
	}
	return nil
}
