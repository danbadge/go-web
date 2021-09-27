package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

type Response struct {
	Message string ` json:"message" `
}

func (c App) Ping() revel.Result {
	result := Response{Message: "pong"}
	return c.RenderJSON(result)
}