package controllers

import "log"

type AdminController struct {
	l *log.Logger
}

func NewAdminController(l *log.Logger) *AdminController {
	return &AdminController{l}
}
