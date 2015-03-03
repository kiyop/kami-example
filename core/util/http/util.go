package http

import (
	net_http "net/http"
	"net/textproto"
)

func HeaderExists(r *net_http.Request, key string) bool {
	_, ok := r.Header[textproto.CanonicalMIMEHeaderKey(key)]
	return ok
}

func GetHeader(r *net_http.Request, key string) string {
	h, ok := r.Header[textproto.CanonicalMIMEHeaderKey(key)]
	if !ok {
		return ""
	} else {
		return h[0]
	}
}

func SetHeader(w net_http.ResponseWriter, key, value string) {
	w.Header().Set(key, value)
}
