package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ilmedova/go-url-shortener/internal/services"
	"net/http"
)

type URLHandler struct {
	Service *services.URLService
}

func NewURLHandler(service *services.URLService) *URLHandler {
	return &URLHandler{Service: service}
}

func (h *URLHandler) ShortenURL(c *gin.Context) {
	var req struct{ URL string }
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	short, err := h.Service.ShortenURL(context.Background(), req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not shorten URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": short})
}

func (h *URLHandler) ResolveURL(c *gin.Context) {
	short := c.Param("short")
	original, err := h.Service.GetOriginalURL(context.Background(), short)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, original)
}
