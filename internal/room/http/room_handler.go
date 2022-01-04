package http

import (
	"dplatform/internal/room/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	RoomService domain.RoomService
}

func NewRoomHandler(rs domain.RoomService) *RoomHandler {
	return &RoomHandler{
		RoomService: rs,
	}
}

func (h *RoomHandler) GetRoom(c *gin.Context) {
	roomId, err := strconv.ParseInt(c.Param("room_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context := c.Request.Context()
	user, err := h.RoomService.GetRoom(context, int(roomId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var room domain.CreateRoomParams
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context := c.Request.Context()
	result, saveErr := h.RoomService.CreateRoom(context, &room)
	if saveErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": saveErr.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
