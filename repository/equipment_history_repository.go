package repository

import (
	"miniproject-3/model/domain"

	"gorm.io/gorm"
)

type EquipmentRentalHistoryRepository interface {
	GetAllHistory() ([]domain.EquipmentRentalHistory, error)
}

type EquipmentRentalHistoryImpl struct {
	DB *gorm.DB
}

func (repo *EquipmentRentalHistoryImpl) GetAllHistory() ([]domain.EquipmentRentalHistory, error) {
	var histories []domain.EquipmentRentalHistory
	if err := repo.DB.Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}
