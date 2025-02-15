package utils

import (
	"errors"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var ErrNoSuchOp = errors.New("no such op")

func Name2Op(eventName string) (op fsnotify.Op, err error) {
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

func Op2Name(op fsnotify.Op) string {
	switch op {
	case fsnotify.Write:
		return "write"
	case fsnotify.Remove:
		return "remove"
	case fsnotify.Rename:
		return "rename"
	case fsnotify.Chmod:
		return "chmod"
	}
	return "unknown"
}
