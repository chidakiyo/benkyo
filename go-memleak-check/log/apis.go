package log

import (
	"context"
	"encoding/json"
	"fmt"
	"go.opencensus.io/trace"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

type contextSkipKey struct{}

func NewRequestLogEntry(ctx context.Context, r *http.Request, severity LogSeverity, status int, responseSize int64, startAt time.Time, prjId string) *LogEntry {
	u := *r.URL
	u.Fragment = ""

	traceID := ""
	spanID := ""

	if span := trace.FromContext(ctx); span != nil {
		// 一般用
		traceID = fmt.Sprintf("projects/%s/traces/%s", prjId, span.SpanContext().TraceID.String())
		spanID = span.SpanContext().SpanID.String()

	} else if traceHeader := r.Header.Get("X-Cloud-Trace-Context"); traceHeader != "" {
		// AppEngine とか用
		ss := strings.SplitN(traceHeader, "/", 2)
		traceID = fmt.Sprintf("projects/%s/traces/%s", prjId, ss[0])

		if len(ss) == 2 {
			ss = strings.SplitN(ss[1], ";", 2)
			spanID = ss[0]
		}
	}

	remoteIP := ""
	if v := r.Header.Get("X-AppEngine-User-IP"); v != "" {
		remoteIP = v
	} else if v := r.Header.Get("X-Forwarded-For"); v != "" {
		remoteIP = v
	} else {
		remoteIP = strings.SplitN(r.RemoteAddr, ":", 2)[0]
	}

	endAt := time.Now()

	falseV := false
	httpRequestEntry := &HttpRequest{
		RequestMethod:                  r.Method,
		RequestURL:                     u.RequestURI(),
		RequestSize:                    r.ContentLength,
		Status:                         status,
		ResponseSize:                   responseSize,
		UserAgent:                      r.UserAgent(),
		RemoteIP:                       remoteIP,
		Referer:                        r.Referer(),
		Latency:                        Duration(endAt.Sub(startAt)),
		CacheLookup:                    &falseV,
		CacheHit:                       &falseV,
		CacheValidatedWithOriginServer: &falseV,
		CacheFillBytes:                 nil,
		Protocol:                       r.Proto,
	}

	operation, ok := ctx.Value(contextOperationKey{}).(*LogEntryOperation)
	if !ok {
		operation = nil
	}

	logEntry := &LogEntry{
		Severity:    severity,
		Time:        Time(endAt),
		HttpRequest: httpRequestEntry,
		Trace:       traceID,
		SpanID:      spanID,
		Operation:   operation,
	}

	return logEntry
}

func NewAppLogEntry(ctx context.Context, severity LogSeverity, prjId string) *LogEntry {
	logger, ok := ctx.Value(contextLoggerKey{}).(RequestLogger)
	if !ok {
		panic("unexpected ctx. use ctx from NewRequestLogger")
	}
	logger.PushSeverity(severity)

	traceID := ""
	spanID := ""

	if span := trace.FromContext(ctx); span != nil {
		// X-Cloud-Trace-Context のケアはOpenCensusレベルで行っておく

		traceID = fmt.Sprintf("projects/%s/traces/%s", prjId, span.SpanContext().TraceID.String())
		spanID = span.SpanContext().SpanID.String()
	}

	operation, ok := ctx.Value(contextOperationKey{}).(*LogEntryOperation)
	if !ok {
		operation = nil
	}

	logEntry := &LogEntry{
		Severity:       severity,
		Time:           Time(time.Now()),
		Trace:          traceID,
		SpanID:         spanID,
		Operation:      operation,
		SourceLocation: NewLogEntrySourceLocation(ctx),
	}

	return logEntry
}

func NewLogEntrySourceLocation(ctx context.Context) *LogEntrySourceLocation {

	skip, ok := ctx.Value(contextSkipKey{}).(int)
	if !ok {
		skip = 2
	}

	var sl *LogEntrySourceLocation
	if pc, file, line, ok := runtime.Caller(skip); ok {
		sl = &LogEntrySourceLocation{
			File: file,
			Line: int64(line),
		}
		if function := runtime.FuncForPC(pc); function != nil {
			sl.Function = function.Name()
		}
	}

	return sl
}

func RequestLog(ctx context.Context, r *http.Request, severity LogSeverity, status int, responseSize int64, startAt time.Time) {
	prjId := os.Getenv("GOOGLE_CLOUD_PROJECT") // FIXME 一旦適当に
	logEntry := NewRequestLogEntry(ctx, r, severity, status, responseSize, startAt, prjId)
	b, err := json.Marshal(logEntry)
	if err != nil {
		panic(err)
	}
	_, _ = fmt.Fprintln(os.Stderr, string(b))
}

func AppLogf(ctx context.Context, severity LogSeverity, format string, a ...interface{}) {

	ctx = context.WithValue(ctx, contextSkipKey{}, 3)

	prjId := os.Getenv("GOOGLE_CLOUD_PROJECT") // FIXME 一旦適当に
	logEntry := NewAppLogEntry(ctx, severity, prjId)
	logEntry.Message = fmt.Sprintf(format, a...)

	b, err := json.Marshal(logEntry)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

