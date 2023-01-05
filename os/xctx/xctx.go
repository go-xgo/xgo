package xctx

import (
	"context"
)

type (
	Ctx    = context.Context // Ctx is short name alias for context.Context.
	StrKey string            // StrKey is a type for warps basic type string as context key.
)

var (
	// initCtx is the context initialized from process environment.
	initCtx context.Context
)

// New creates and returns a context which contains context id.
func New() context.Context {
	return WithCtx(context.Background())
}

// WithCtx creates and returns a context containing context id upon given parent context `ctx`.
func WithCtx(ctx context.Context) context.Context {
	return ctx
}
