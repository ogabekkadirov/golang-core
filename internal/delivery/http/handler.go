package http

import "golang-core/pkg/auth"

type Handler struct {
	tokenManager auth.TokenManager
	
}

func NewHandler()