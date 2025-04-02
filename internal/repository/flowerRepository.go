package repository

import (
	"FlowerShop/internal/models"
	"gorm.io/gorm"
)

type FlowerRepositoryImpl struct {
	db *gorm.DB
}

// NewStudentRepository - Constructor
func NewFlowerRepository(db *gorm.DB) *FlowerRepositoryImpl {
	return &FlowerRepositoryImpl{db: db}
}

// GetAll - Retrieve all students
func (f FlowerRepositoryImpl) GetAll() ([]models.Flower, error) {
	var flowers []models.Flower
	err := f.db.Find(&flowers).Error
	return flowers, err
}

// GetById - Retrieve a student by ID
func (f FlowerRepositoryImpl) GetById(id int) (*models.Flower, error) {
	var flower models.Flower
	err := f.db.First(&flower, id).Error
	return &flower, err
}

// Create - Add a new student
func (f FlowerRepositoryImpl) Create(flower *models.Flower) error {
	return f.db.Create(flower).Error
}

// Update - Modify an existing student
func (f FlowerRepositoryImpl) Update(id int, flower *models.FlowerEdit) error {
	return f.db.Model(&models.Flower{}).Where("id = ?", id).Omit("id, CreatedAt").Updates(flower).Error
}

// Delete - Remove a student by ID
func (f FlowerRepositoryImpl) Delete(flowerID int) error {
	return f.db.Delete(&models.Flower{}, flowerID).Error
}
