package repositories

import (
	"database/sql"

	"github.com/denisemignoli/to-do-list/models"
)

type UserMySQLRepository struct {
	db *sql.DB
}

func NewUserMySQLRepository(db *sql.DB) *UserMySQLRepository {
	return &UserMySQLRepository{
		db: db,
	}
}

func (ur *UserMySQLRepository) CreateUser(user models.User) (int64, error) {
	result, err := ur.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ur *UserMySQLRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := ur.db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
