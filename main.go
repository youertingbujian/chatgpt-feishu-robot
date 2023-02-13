// Start a web server
package main

import (
	"chatgpt/biz"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1Api := api.Group("/v1")
		v1Api.POST("/webhook/event", biz.ReceiveEvent)

		//v2Api := api.Group("v2")
		//v2Api.GET("/demo", biz.demo)
	}

	if err := r.Run(":51515"); err != nil {
		logrus.WithError(err).Errorf("init fail")
	}
}
