package utils

import (
	"errors"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var ErrNoSuchOp = errors.New("no such op")

func OpByName(eventName string) (op fsnotify.Op, err error) {
	switch strings.ToLower(eventName) {
	case "write":
		op = fsnotify.Write
	case "remove":
		op = fsnotify.Remove
	case "rename":
		op = fsnotify.Rename
	case "chmod":
		op = fsnotify.Chmod
	default:
		err = ErrNoSuchOp
	}
	return op, err
}
