package repositories

import (
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/models"
	"gorm.io/gorm"
)


type PaymentRepository struct {
	Db *gorm.DB
}

func (pym *PaymentRepository) CreatePayment(payment *models.Payment) error {
	return pym.Db.Create(payment).Error
}