package main

import (
	"example/internal/application"
	"example/internal/infra/redisimpl"
	"example/pkg/circuitbreaker"
	"example/pkg/service1"
	"example/pkg/service2"

	"github.com/gin-gonic/gin"
)

func main() {
	service1breaker := circuitbreaker.New(redisimpl.NewRedisImpl(), "service1", 1)
	s1 := service1.New(service1breaker)

	service2breaker := circuitbreaker.New(redisimpl.NewRedisImpl(), "service2", 2)
	s2 := service2.New(service2breaker)

	app := application.New(s1, s2)

	r := gin.Default()

	r.GET("/dosomething/service1", func(c *gin.Context) {
		c.JSON(200, app.Service1())
	})

	r.GET("/dosomething/service2", func(c *gin.Context) {
		c.JSON(200, app.Service2())
	})

	r.GET("/state/service1", func(c *gin.Context) {
		c.JSON(200, service1breaker.State())
	})

	r.GET("/state/service2", func(c *gin.Context) {
		c.JSON(200, service2breaker.State())
	})

	r.Run()
}
