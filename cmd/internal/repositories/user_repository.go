package repositories

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

// Method to create a new user
func (ur *UserRepository) CreateUser(user *models.User) error{
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

// Method to update user
func (ur *UserRepository) UpdateUser(user *models.User) error {
    return ur.Db.Save(user).Error
}

// Method to delete user
func (ur *UserRepository) DeleteUser(user *models.User) error {
    return ur.Db.Delete(user).Error
}

