package repositories

import (
	"database/sql"
	"errors"

	"github.com/kerimcetinbas/goginpostgrestut/database"
	"github.com/kerimcetinbas/goginpostgrestut/types"
)

type userRepository struct {
	db *sql.DB
}

type IUserRepository interface {
	CreateUser(data *types.UserCreateDto) error
	GetUsers() (*[]types.User, error)
	FindUserByName(name string) (types.User, error)
	FindUserById(id uint) (types.User, error)
}

func UserRepository() IUserRepository {
	return &userRepository{
		db: database.DB,
	}
}

func (r *userRepository) CreateUser(data *types.UserCreateDto) error {
	var (
		err error
	)
	_, err = r.db.Exec(`
	 INSERT INTO users (name, password) VALUES ($1, $2)
	`, data.Name, data.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUsers() (*[]types.User, error) {
	users := make([]types.User, 0)
	rows, err := r.db.Query(`
	SELECT u.id as id, u.name as user, COUNT(m.message) as messages
    FROM users  u
    LEFT JOIN messages m on u.id = m.user_id
    GROUP BY u.id
	`)

	if err != nil {
		return &users, errors.New("could not find")
	}
	for rows.Next() {
		user := types.User{}
		rows.Scan(&user.ID, &user.Name, &user.Messages)
		users = append(users, user)
	}
	return &users, nil
}

func (r *userRepository) FindUserById(id uint) (types.User, error) {
	var (
		row  *sql.Rows
		err  error
		user types.User
	)
	row, err = r.db.Query(`
		SELECT id, name
		FROM users
		WHERE id = $1
	`, id)

	if err != nil {
		return user, err
	}
	if row.Next() {
		err = row.Scan(&user.ID, &user.Name)

		if err != nil {
			return user, err
		}
	}

	return user, nil
}
func (r *userRepository) FindUserByName(name string) (types.User, error) {
	var (
		row  *sql.Rows
		err  error
		user types.User
	)

	row, err = r.db.Query(`
		SELECT id, name, password FROM users 
		WHERE name=$1
		LIMIT 1
	`, name)

	if err != nil {
		return user, errors.New("User not found")
	}

	if row.Next() {
		err = row.Scan(&user.ID, &user.Name, &user.Password)
		if err != nil {
			return user, errors.New("User not found")
		}
		return user, nil
	}

	return user, errors.New("User not found")

}
