package middleware

import (
	"net/http"
	"strings"
)

type RemoteAddressMiddleware struct {
}

func NewRemoteAddressMiddleware() *RemoteAddressMiddleware {
	return &RemoteAddressMiddleware{}
}

func (m *RemoteAddressMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var addr string
		addrs, ok := r.Header["X-Forwarded-For"]
		if !ok {
			addr = r.RemoteAddr
		} else {
			addr = addrs[0]
		}
		// extract ip
		ip := strings.Split(addr, ":")[0]
		r.Header["Remote-Address"] = []string{ip}

		// Passthrough to next handler if need
		next(w, r)
	}
}
