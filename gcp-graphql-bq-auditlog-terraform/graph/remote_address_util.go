package graph

import (
	"context"
	"go.opencensus.io/trace"
	"net/http"
	"strings"
)

type RequestMeta struct {
	RemoteIP string
	TraceID  string
	SpanID   string
}
type contextKey struct {
	name string
}

var requestCtxKey = &contextKey{"request"}

func RequestMetaFromContext(ctx context.Context) *RequestMeta {
	raw, _ := ctx.Value(requestCtxKey).(*RequestMeta)
	return raw
}

func NewRequestMetaContext(ctx context.Context, requestCtx *RequestMeta) context.Context {
	return context.WithValue(ctx, requestCtxKey, requestCtx)
}

func RequestMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sc := trace.FromContext(r.Context()).SpanContext()
			ctx := NewRequestMetaContext(r.Context(), &RequestMeta{
				RemoteIP: getRemoteAddr(r),
				TraceID:  sc.TraceID.String(),
				SpanID:   sc.SpanID.String(),
			})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getRemoteAddr(r *http.Request) string {
	addr := r.Header.Get("X-Real-Ip")
	if addr == "" {
		xff := r.Header.Get("X-Forwarded-For")
		if strings.Contains(xff, ",") {
			addr = xff[:strings.Index(xff, ",")]
		}
	}
	if addr == "" {
		if strings.Contains(r.RemoteAddr, ":") {
			addr = r.RemoteAddr[:strings.Index(r.RemoteAddr, ":")]
		} else {
			addr = r.RemoteAddr
		}
	}
	return addr
}
