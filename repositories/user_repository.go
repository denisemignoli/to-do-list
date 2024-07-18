package repositories

import "github.com/denisemignoli/to-do-list/models"

type UserRepository interface {
	CreateUser(user models.User) (int64, error)
	GetUserByUsername(username string) (*models.User, error)
}
