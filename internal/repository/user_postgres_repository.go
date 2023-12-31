package repository

import (
	"database/sql"
	"golang-project-template/internal/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) (int, error) {
	var id int
	row := r.db.QueryRow("INSERT into USERS (name, email) VALUES ($1, $2)", user.Name, user.Email)

	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}

func (r *userRepository) GetByID(id int) (*models.User, error) {
	row := r.db.QueryRow("select * from users where id = $1", id)

	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]*models.User, error) {
	rows, err := r.db.Query("select * from users")
	if err != nil {
		return nil, err
	}

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)

	}

	return users, nil
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	_, err := r.db.Exec("UPDATE users SET name=$1, email=$2 where id=$3", user.Name, user.Email, user.Id)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow("Select * from users where id=$1", user.Id).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE from users where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
