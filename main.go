package main

import (
	"example/internal/application"
	"example/internal/infra/breakerstorage"
	"example/pkg/circuitbreaker"
	"example/pkg/service1"
	"example/pkg/service2"
	"time"

	"github.com/andreleoni/testpkggolang"
	"github.com/gin-gonic/gin"
)

func main() {
	service1breaker := circuitbreaker.New(breakerstorage.NewRedisImpl(), "service1", 1, 30*time.Second, 5)
	s1 := service1.New(service1breaker)

	service2breaker := circuitbreaker.New(breakerstorage.NewRedisImpl(), "service2", 2, 30*time.Second, 5)
	s2 := service2.New(service2breaker)

	app := application.New(s1, s2)

	r := gin.Default()

	r.GET("/service1", func(c *gin.Context) {
		c.JSON(200, app.Service1())
	})

	r.GET("/service2", func(c *gin.Context) {
		c.JSON(200, app.Service2())
	})

	testpkggolang.Start()

	r.Run()
}
