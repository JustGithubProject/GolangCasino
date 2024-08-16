package repositories

import (
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/models"
	"gorm.io/gorm"
)


type PaymentRepository struct {
	Db *gorm.DB
}

func (pym *PaymentRepository) CreatePayment(payment *models.Payment) error {
	/*
		Creating payment
	*/
	return pym.Db.Create(payment).Error
}

func (pym *PaymentRepository) GetPaymentByOrderID(orderID string) (*models.Payment, error){
	/*
		Getting payment by ORDER_ID
	*/
	var payment models.Payment
	err := pym.Db.First(&payment, "order_id = ?", orderID).Error
	if err != nil{
		return nil, err
	}
	return &payment, nil
}


func (pym *PaymentRepository) UpdateStatusPayment(payment *models.Payment) error {
	return pym.Db.Model(payment).Update("Status", payment.Status).Error
}