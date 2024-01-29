package handlers

import "github.com/blezzio/triniti/apis/interfaces"

type UrlOpt func(h *URL)

func WithView(v interfaces.View, name ViewName) UrlOpt {
	return func(h *URL) {
		h.views[name] = v
	}
}
