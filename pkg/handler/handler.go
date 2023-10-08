package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// GetOrder godoc
// @Summary Get Oder By Id
// @Description Responds with the order as JSON
// @Param    id  path      int  true  "Search order by ID"
// @Produce  json
// @Success 200 {object} wb_l0.Order
// @Failure 400,404 {object} error
// @Router /{id} [get]
func (h *Handler) getOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))
	fmt.Println(orderId)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	order, err := h.service.GetOrder(orderId)
	if err != nil {
		newErrorResponce(c, http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, order)
	}
}
