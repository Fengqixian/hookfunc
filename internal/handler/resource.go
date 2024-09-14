package handler

import (
	"github.com/gin-gonic/gin"
	"hookfunc/internal/service"
)

type ResourceHandler struct {
	*Handler
	resourceService service.ResourceService
}

func NewResourceHandler(handler *Handler, resourceService service.ResourceService) *ResourceHandler {
	return &ResourceHandler{
		Handler:         handler,
		resourceService: resourceService,
	}
}

func (h *ResourceHandler) GetResource(ctx *gin.Context) {

}
