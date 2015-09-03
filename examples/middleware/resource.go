package main

import (
	"github.com/leonelquinteros/yarf"
)

type Hello struct {
	yarf.Resource
}

// Implement the GET handler with optional name parameter
func (h *Hello) Get() error {
	name := h.Context.Params.Get("name")

	salute := "Hello"
	if name != "" {
		salute += ", " + name
	}
	salute += "!"

	h.Render(salute)

	return nil
}