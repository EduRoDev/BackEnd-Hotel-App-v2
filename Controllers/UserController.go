package controllers

import "log"

type UserController struct {
	l *log.Logger
}

func NewUserController(l *log.Logger) *UserController {
	return &UserController{l}
}
