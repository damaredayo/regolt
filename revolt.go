package regolt

import "net/http"

const VERSION = "v0.1.0"

func CreateBot(token string) *Session {
	return &Session{
		Token:    token,
		handlers: make(map[string][]eventHandler),
		http:     http.DefaultClient,
	}
}
