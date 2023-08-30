package handler

import (
	"avito-dynamic-segment-back/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}

		segments := users.Group(":id/segments")
		{
			segments.POST("/", h.createSegment)
			segments.GET("/", h.getAllSegments)
			segments.GET("/:segment_id", h.getSegmentById)
			segments.PUT("/:segment_id", h.updateSegment)
			segments.DELETE("/:segment_id", h.deleteSegment)
		}
	}

	return router
}
