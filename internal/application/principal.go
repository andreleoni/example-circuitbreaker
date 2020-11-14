package application

import (
	"example/pkg/anotherservice"
	"example/pkg/oneservice"
)

type Application struct {
	oneService     *oneservice.OneService
	anotherService *anotherservice.AnotherService
}

func New(oneService *oneservice.OneService, anotherService *anotherservice.AnotherService) *Application {
	return &Application{oneService, anotherService}
}

func (app Application) Principal() {
	app.oneService.DoSomething()

	app.anotherService.CreateOrder()
	app.anotherService.CreateOrder()
	app.anotherService.CreateOrder()
	app.anotherService.CreateOrder()
}
