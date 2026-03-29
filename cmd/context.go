package cmd

import (
	"context"
	"io"
	"os"
)

type contextKey string

const contextObjectKey = contextKey("contextObject")

type contextObject struct {
	File *os.File
	Out  io.Writer
}

func createContext(file *os.File, out io.Writer) context.Context {
	return context.WithValue(context.Background(), contextObjectKey, &contextObject{
		File: file,
		Out:  out,
	})
}

func ContextObjectFromContext(ctx context.Context) (*contextObject, bool) {
	obj, ok := ctx.Value(contextObjectKey).(*contextObject)
	return obj, ok
}
