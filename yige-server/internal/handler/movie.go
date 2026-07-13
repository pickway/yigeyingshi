package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-server/internal/model"
	"github.com/yigeyingshi/yige-server/internal/service"
)

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(s *service.MovieService) *MovieHandler {
	return &MovieHandler{service: s}
}

func (h *MovieHandler) List(c *gin.Context) {
	genre := c.Query("genre")
	sort := c.DefaultQuery("sort", "rating_desc")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "12"))

	list, total, err := h.service.List(genre, sort, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  list,
		"total": total,
		"page":  page,
		"pageSize": pageSize,
	})
}

func (h *MovieHandler) Featured(c *gin.Context) {
	list, err := h.service.Featured(3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *MovieHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	movie, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func (h *MovieHandler) Create(c *gin.Context) {
	var m model.Movie
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": m})
}
