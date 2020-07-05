package controllers

import "review-server/core/http/request"

type ReController struct {
}

func (r *ReController) Request() request.RequestBodyInterface {
	return request.GetRequestBody()
}
