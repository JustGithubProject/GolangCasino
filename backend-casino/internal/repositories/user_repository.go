package repositories

import (
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

// Method to create a new user
func (ur *UserRepository) CreateUser(user *models.User) error{
	return ur.Db.Create(user).Error
}

func (ur *UserRepository) CreateGoogleUser(user* models.User) error {
	return ur.Db.Create(user).Error
}


// Method to get user by id
func (ur *UserRepository) GetUserById(id uint) (*models.User, error){
	var user models.User
	err := ur.Db.First(&user, id).Error
	if err != nil{
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByGoogleID(googleID string) (*models.User, error) {
	var user models.User
	// err := ur.Db.First(&user, "GoogleID = ?", googleID).Error
	err := ur.Db.First(&user, "google_id = ?", googleID).Error
	if err != nil {
		return nil, err 
	}
	return &user, nil
}


// Method to get user by id
func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error){
	var user models.User
	err := ur.Db.First(&user, "Name = ?", username).Error
	if err != nil{
		return nil, err
	}
	return &user, nil
}

// Method to update user
func (ur *UserRepository) UpdateUser(user *models.User) error {
    return ur.Db.Save(user).Error
}

// Method to delete user
func (ur *UserRepository) DeleteUser(user *models.User) error {
    return ur.Db.Delete(user).Error
}

func (ur *UserRepository) UpdateBalanceUser(user *models.User) error {
    return ur.Db.Model(user).Update("balance", user.Balance).Error
}

func (ur *UserRepository) GetUserPayments(userID uint) (*models.User, error) {
    var user models.User

    // Find the user by ID and preload payments
    result := ur.Db.Preload("Payments").First(&user, userID)
    if result.Error != nil {
        return nil, result.Error // Return the error if the user or payments cannot be fetched
    }

    return &user, nil // Return the user with preloaded payments
}