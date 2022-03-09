package tracer

import (
	"context"
)

type Span struct {
}

type spanWithContextOptions struct {
	environment   string
	operationName string
	logErrors     bool
	logText       string
}

type spanWithContextOptionModifier func(opts *spanWithContextOptions)

func (s Span) SetTag(tagName string, value interface{}) {
	return
}

type finishCallback func(errors ...error)

func finish(errors ...error) {
	return
}

// Trace creates and return new span from passed context
// Common usage:
// var err error
// ctx, span, finish := tracer.Trace(ctx)
// span.SetTag(tag.TagName, value)
// defer func(){ finish(err) }()
// Options:
//   WithErrorLogging(logText string)        - additionally writes error to log
//   WithEnvironment(environment string)     - sets custom environment
//   WithOperationName(operationName string) - sets custom span operation name
//   ExtractFrom(header http.Header)         - get trace data from given header
func Trace(ctx context.Context, opts ...spanWithContextOptionModifier) (
	context.Context,
	Span,
	finishCallback,
) {
	return ctx, Span{}, finish
}
