package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wb_l0/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/:id", h.getOrder)

	return router
}

func (h *Handler) getOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Invalid ID")
	}

	order, err := h.service.GetOrder(orderId)
	if err != nil {
		newErrorResponce(c, http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, order)
	}
}
