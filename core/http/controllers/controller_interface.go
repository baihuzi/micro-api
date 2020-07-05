package controllers

import "review-server/core/http/request"

type ControllerInterface interface {
	Request() request.RequestBodyInterface
}
