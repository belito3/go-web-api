package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// Define key name
const (
	TraceIDKey      = "trace_id"
	UserIDKey       = "user_id"
	SpanTitleKey    = "span_title"
	SpanFunctionKey = "span_function"
	VersionKey      = "version"
	StackKey        = "stack"
)

// TraceIDFunc Define the function to get the tracking ID
type TraceIDFunc func() string

var (
	version     string
	traceIDFunc TraceIDFunc
	pid         = os.Getpid()
)

func init() {
	traceIDFunc = func() string {
		//return fmt.Sprintf("trace-id-%d-%s",
		//	os.Getpid(),
		//	time.Now().Format("2006.01.02.15.04.05.999999"))
		return fmt.Sprintf("trace-id-%d",
			os.Getpid())
	}
}

// Logger Define log alias
type Logger = logrus.Logger

// Hook Define log hook alias
type Hook = logrus.Hook

// StandardLogger Get standard log
func StandardLogger() *Logger {
	return logrus.StandardLogger()
}

// SetLevel Set log level
func SetLevel(level int) {
	logrus.SetLevel(logrus.Level(level))
}

// SetFormatter Set log output format
func SetFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(new(logrus.JSONFormatter))
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			//DisableColors: true,
			FullTimestamp: true,
			//TimestampFormat: time.RFC3339Nano,
		})
	}
}

// SetOutput Set log output
func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

// SetVersion Set version
func SetVersion(v string) {
	version = v
}

// SetTraceIDFunc Set the processing function of the tracking ID
func SetTraceIDFunc(fn TraceIDFunc) {
	traceIDFunc = fn
}

// AddHook Add log hook
func AddHook(hook Hook) {
	logrus.AddHook(hook)
}

type (
	traceIDKey struct{}
	userIDKey  struct{}
)

// NewTraceIDContext Create tracking ID context
func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// FromTraceIDContext Get the tracking ID from the context
func FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return traceIDFunc()
}

// NewUserIDContext Create user ID context
func NewUserIDContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

// FromUserIDContext Get user ID from context
func FromUserIDContext(ctx context.Context) string {
	v := ctx.Value(userIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

type spanOptions struct {
	Title    string
	FuncName string
}

// SpanOption Define the data items of the tracking unit
type SpanOption func(*spanOptions)

// SetSpanTitle Set the title of the tracking unit
func SetSpanTitle(title string) SpanOption {
	return func(o *spanOptions) {
		o.Title = title
	}
}

// SetSpanFuncName Set the function name of the tracking unit
func SetSpanFuncName(funcName string) SpanOption {
	return func(o *spanOptions) {
		o.FuncName = funcName
	}
}

// StartSpan Start a tracking unit
func StartSpan(ctx context.Context, opts ...SpanOption) *Entry {
	if ctx == nil {
		ctx = context.Background()
	}

	var o spanOptions
	for _, opt := range opts {
		opt(&o)
	}

	fields := map[string]interface{}{
		VersionKey: version,
	}
	if v := FromTraceIDContext(ctx); v != "" {
		fields[TraceIDKey] = v
	}
	if v := FromUserIDContext(ctx); v != "" {
		fields[UserIDKey] = v
	}
	if v := o.Title; v != "" {
		fields[SpanTitleKey] = v
	}
	if v := o.FuncName; v != "" {
		fields[SpanFunctionKey] = v
	}

	return newEntry(logrus.WithFields(fields))
}

// Debugf Write debug log
func Debugf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Debugf(format, args...)
}

// Infof Write message log
func Infof(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Infof(format, args...)
}

// Printf Write message log
func Printf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Printf(format, args...)
}

// Warnf Write warning log
func Warnf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Warnf(format, args...)
}

// Errorf Write error log
func Errorf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Errorf(format, args...)
}

// Fatalf Write to major error log
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Fatalf(format, args...)
}

// ErrorStack Output error stack
func ErrorStack(ctx context.Context, err error) {
	StartSpan(ctx).WithField(StackKey, fmt.Sprintf("%+v", err)).Errorf(err.Error())
}

func newEntry(entry *logrus.Entry) *Entry {
	return &Entry{entry: entry}
}

// Entry Define a unified log writing method
type Entry struct {
	entry *logrus.Entry
}

func (e *Entry) checkAndDelete(fields map[string]interface{}, keys ...string) {
	for _, key := range keys {
		_, ok := fields[key]
		if ok {
			delete(fields, key)
		}
	}
}

// WithFields Structured field write
func (e *Entry) WithFields(fields map[string]interface{}) *Entry {
	e.checkAndDelete(fields,
		TraceIDKey,
		SpanTitleKey,
		SpanFunctionKey,
		VersionKey)
	return newEntry(e.entry.WithFields(fields))
}

// WithField Structured field write
func (e *Entry) WithField(key string, value interface{}) *Entry {
	return e.WithFields(map[string]interface{}{key: value})
}

// Fatalf Major error log
func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.entry.Fatalf(format, args...)
}

// Errorf Error log
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.entry.Errorf(format, args...)
}

// Warnf Warning log
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.entry.Warnf(format, args...)
}

// Infof Message log
func (e *Entry) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

// Printf Message log
func (e *Entry) Printf(format string, args ...interface{}) {
	e.entry.Printf(format, args...)
}

// Debugf Write debug log
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.entry.Debugf(format, args...)
}