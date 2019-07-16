package main

import (
	"github.com/erikdubbelboer/fasthttp"
	routing "github.com/jackwhelpton/fasthttp-routing"
	"github.com/jackwhelpton/fasthttp-routing/access"
	"github.com/jackwhelpton/fasthttp-routing/content"
	"github.com/jackwhelpton/fasthttp-routing/fault"
	"github.com/jackwhelpton/fasthttp-routing/slash"
	"log"
	"sync"
)

type router struct{}

type IRouter interface {
	InitRouter() *routing.Router
}

func (router *router) InitRouter() *routing.Router {
	api := routing.New()

	api.Use(
		// all these handlers are shared by every route
		access.Logger(log.Printf),
		slash.Remover(fasthttp.StatusMovedPermanently),
		fault.Recovery(log.Printf),
	)

	api.Use(
		// these handlers are shared by the routes
		content.TypeNegotiator(content.JSON),
	)

	categoriesController := ServiceContainer().InjectCategoriesController()

	api.Get("/categories/<categoryCode>", func(c *routing.Context) error {
		return categoriesController.GetCategoryWithProducts(c)
	})

	return api
}

var (
	m          *router
	routerOnce sync.Once
)

func Router() IRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
