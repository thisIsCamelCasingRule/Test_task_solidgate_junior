package api

import "Test_task_solidgate_junior/internal/service"

type Api struct {
	srv service.Serve
}

func NewApi() Api {
	return Api{
		srv: service.NewService(),
	}
}

type Handle interface {
	CardHandler
}
