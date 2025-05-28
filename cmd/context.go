package cmd

import (
	"context"
	"os"
)

type contextKey string

const fileContextKey contextKey = "excludeFile"

// ContextWithFile returns a new context with the given file set.
func ContextWithFile(ctx context.Context, file *os.File) context.Context {
	return context.WithValue(ctx, fileContextKey, file)
}

// FileFromContext retrieves the file from the context, if it exists.
func FileFromContext(ctx context.Context) (*os.File, bool) {
	file, ok := ctx.Value(fileContextKey).(*os.File)
	return file, ok
}
