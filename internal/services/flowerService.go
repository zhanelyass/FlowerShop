package services

import (
	"FlowerShop/internal/models"
)

type FlowerRepository interface {
	GetAll() ([]models.Flower, error)
	GetById(id int) (*models.Flower, error)
	Create(flower *models.Flower) error
	Update(id int, flower *models.FlowerEdit) error
	Delete(flowerID int) error
}

type FlowerService struct {
	repo FlowerRepository
}

// Конструктор StudentService
func NewFlowerService(flowerRepo FlowerRepository) *FlowerService {
	return &FlowerService{repo: flowerRepo}
}

// Получение всех студентов
func (f *FlowerService) GetAllFlowers() ([]models.Flower, error) {
	return f.repo.GetAll()
}

// Получение студента по ID
func (f *FlowerService) GetFlowerByID(id int) (*models.Flower, error) {
	return f.repo.GetById(id)
}

// Создание нового студента
func (f *FlowerService) Create(name, description string, price float64) (*models.Flower, error) {
	flower := &models.Flower{
		Name:        name,
		Description: description,
		Price:       price,
	}
	err := f.repo.Create(flower)
	return flower, err
}

// Обновление данных студента
func (f *FlowerService) Update(id int, flowerEdit *models.FlowerEdit) (*models.Flower, error) {
	err := f.repo.Update(id, flowerEdit)
	if err != nil {
		return nil, err
	}
	return f.GetFlowerByID(id)
}

// Удаление студента
func (f *FlowerService) DeleteFlower(flowerID int) error {
	return f.repo.Delete(flowerID)
}
