package repositories

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// Method to create a new user
func (ur *UserRepository) CreateUser(user *models.User) error{
	return ur.db.Create(user).Error
}


// Method to get user by id
func (ur *UserRepository) GetUserById(id uint) (*models.User, error){
	var user models.User
	err := ur.db.First(&user, id).Error
	if err != nil{
		return nil, err
	}
	return &user, nil
}

// Method to update user
func (ur *UserRepository) UpdateUser(user *models.User) error {
    return ur.db.Save(user).Error
}

// Method to delete user
func (ur *UserRepository) DeleteUser(user *models.User) error {
    return ur.db.Delete(user).Error
}

