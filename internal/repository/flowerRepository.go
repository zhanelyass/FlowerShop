package repository

import (
	"FlowerShop/internal/models"
	"gorm.io/gorm"
)

type FlowerRepositoryImpl struct {
	db *gorm.DB
}

// Конструктор
func NewFlowerRepository(db *gorm.DB) *FlowerRepositoryImpl {
	return &FlowerRepositoryImpl{db: db}
}

// Получить все цветы
func (f FlowerRepositoryImpl) GetAll() ([]models.Flower, error) {
	var flowers []models.Flower
	err := f.db.Find(&flowers).Error
	return flowers, err
}

// Получить цветок по ID
func (f FlowerRepositoryImpl) GetById(id int) (*models.Flower, error) {
	var flower models.Flower
	err := f.db.First(&flower, id).Error
	return &flower, err
}

// Создать новый цветок
func (f FlowerRepositoryImpl) Create(flower *models.Flower) error {
	return f.db.Create(flower).Error
}

// Обновить цветок
func (f FlowerRepositoryImpl) Update(id int, flower *models.FlowerEdit) error {
	return f.db.Model(&models.Flower{}).Where("id = ?", id).Omit("id, CreatedAt").Updates(flower).Error
}

// Удалить цветок
func (f FlowerRepositoryImpl) Delete(flowerID int) error {
	return f.db.Delete(&models.Flower{}, flowerID).Error
}
