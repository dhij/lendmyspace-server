package http

import (
	"lendmyspace-server/internal/space/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SpaceHandler struct {
	SpaceService domain.SpaceService
}

func NewSpaceHandler(ss domain.SpaceService) *SpaceHandler {
	return &SpaceHandler{
		SpaceService: ss,
	}
}

func (h *SpaceHandler) GetSpace(c *gin.Context) {
	roomId, err := strconv.ParseInt(c.Param("space_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context := c.Request.Context()
	space, err := h.SpaceService.GetSpace(context, int(roomId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, space)
}

func (h *SpaceHandler) CreateSpace(c *gin.Context) {
	var space domain.CreateSpaceParams
	if err := c.ShouldBindJSON(&space); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context := c.Request.Context()
	result, saveErr := h.SpaceService.CreateSpace(context, &space)
	if saveErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": saveErr.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
