package repositories

import (
	"strings"

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
		db = db.Where("phone_number LIKE ?", "%"+condition["phone_number"].(string))
	}

	if err := db.Find(&userList).Error; err != nil {
		return nil, app.ThrowDefaultNotFoundError(err)
	}

	// Filter full name with userList like a search engine
	if condition["full_name"] != nil && condition["full_name"].(string) != "" {
		var userListFiltered []models.User

		for _, user := range userList {
			fullName := strings.ToLower(condition["full_name"].(string))
			if strings.Contains(strings.ToLower(user.FullName), fullName) {
				userListFiltered = append(userListFiltered, user)
			}
		}

		userList = userListFiltered

		if len(userList) == 0 {
			return nil, app.ThrowDefaultNotFoundError(nil)
		}
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

func (r UserRepository) UpdateUser(condition map[string]any, updateData *models.UserUpdate) (uint, *app.AppError) {
	db := database.GetDB()

	// Create and get user id
	if err := db.Where(condition).Updates(&updateData).Error; err != nil {
		return 0, app.ThrowInternalServerError(err)
	}

	userUpdated, err := r.GetUser(condition)

	if err != nil {
		return 0, err
	}

	return userUpdated.ID, nil
}
