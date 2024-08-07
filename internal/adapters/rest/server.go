package rest

import (
	"fmt"
	"github.com/AnwarSaginbai/todo-list/internal/ports"
	"github.com/gin-gonic/gin"
)

type Adapter struct {
	service ports.APIPort
}

func NewAdapter(service ports.APIPort) *Adapter {
	return &Adapter{service: service}
}

func (a *Adapter) Run(addr string) error {
	r := gin.New()
	r.POST("api/v1/todo", a.PostTaskHandler)
	r.GET("api/v1/todo", a.GetAllHandler)
	r.DELETE("api/v1/todo/:id", a.DeleteTaskHandler)
	r.PATCH("api/v1/todo/:id", a.UpdateTaskHandler)
	if err := r.Run(fmt.Sprintf(":%s", addr)); err != nil {
		return err
	}
	return nil
}
