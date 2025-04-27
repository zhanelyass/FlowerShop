package delivery

import (
	"FlowerShop/internal/models"
	"FlowerShop/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FlowerHandler struct {
	service *services.FlowerService
}

// Конструктор
func NewFlowerHandler(service *services.FlowerService) *FlowerHandler {
	return &FlowerHandler{service: service}
}

// Получение списка всех цветов
func (h *FlowerHandler) GetAllFlowers(c *gin.Context) {
	flowers, err := h.service.GetAllFlowers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flowers"})
		return
	}
	c.JSON(http.StatusOK, flowers)
}

// Получение цветка по ID
func (h *FlowerHandler) GetFlower(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flower ID"})
		return
	}

	flower, err := h.service.GetFlowerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
		return
	}

	c.JSON(http.StatusOK, flower)
}

// Создание нового цветка
func (h *FlowerHandler) CreateFlower(c *gin.Context) {
	var flowerCreate models.FlowerEdit

	if err := c.ShouldBindJSON(&flowerCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newFlower := &models.Flower{
		Name:        flowerCreate.Name,
		Description: flowerCreate.Description,
		Price:       flowerCreate.Price,
	}

	newFlower, err := h.service.Create(newFlower)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flower"})
		return
	}

	c.JSON(http.StatusCreated, newFlower)
}

// Обновление данных цветка
func (h *FlowerHandler) UpdateFlower(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flower ID"})
		return
	}

	var flowerEdit models.FlowerEdit
	if err := c.ShouldBindJSON(&flowerEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedFlower, err := h.service.Update(id, &flowerEdit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
		return
	}

	c.JSON(http.StatusOK, updatedFlower)
}

// Удаление цветка
func (h *FlowerHandler) DeleteFlower(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid flower ID"})
		return
	}

	if err := h.service.DeleteFlower(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flower deleted successfully"})
}
