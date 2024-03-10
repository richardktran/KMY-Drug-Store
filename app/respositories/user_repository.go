package repositories

import (
	"github.com/richardktran/KMY-Drug-Store/app/models"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
	"github.com/richardktran/KMY-Drug-Store/pkg/database"
)

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (r UserRepository) GetUser(condition map[string]any) (*models.User, *app.AppError) {
	var user models.User
	db := database.GetDB()

	if err := db.Where(condition).First(&user).Error; err != nil {
		return nil, app.ThrowDefaultNotFoundError(err)
	}

	return &user, nil
}

func (r UserRepository) GetUserList(condition map[string]any) ([]models.User, *app.AppError) {
	var userList []models.User
	db := database.GetDB()

	if condition["full_name"] != nil && condition["full_name"].(string) != "" {
		db = db.Where("full_name LIKE ?", "%"+condition["full_name"].(string)+"%")
	}

	if condition["phone_number"] != nil && condition["phone_number"].(string) != "" {
		db = db.Where("phone_number LIKE ?", "%"+condition["phone_number"].(string)+"%")
	}

	if err := db.Find(&userList).Error; err != nil {
		return nil, app.ThrowDefaultNotFoundError(err)
	}

	return userList, nil
}

func (r UserRepository) CreateUser(data *models.UserCreation) (uint, *app.AppError) {
	db := database.GetDB()

	// Create and get user id
	if err := db.Create(&data).Error; err != nil {
		return 0, app.ThrowInternalServerError(err)
	}

	return data.ID, nil
}
