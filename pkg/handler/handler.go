package handler

import (
	"github.com/EltIsma/YandexLavka/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()
	rateLimiter := middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(2))
	router.Use(rateLimiter)

	courier := router.Group("/couriers")
	{
		courier.POST("/", h.createCouriers)
		courier.GET("/", h.getAllCouriers)
		courier.GET("/:id", h.getCourierById)
		courier.GET("/meta-info/:id", h.getCourierRatingSalary)
	}

	orders := router.Group("/orders")
	{
		orders.POST("/", h.createOrders)
		orders.GET("/", h.getAllOrders)
		orders.GET("/:id", h.getOrderById)
		orders.POST("/complete", h.ordersComplete)
		//	orders.POST("/assign", h.getAssign)
	}

	return router

}
