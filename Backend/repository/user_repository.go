package repository

import (
	"database/sql"
	"inventory_backend/model"
)

type UserRepository interface {
	GetByID(id int) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(id int, user model.User) (model.User, error)
	Delete(id int) error
	GetByEmail(email string) (model.User, error)
}

// userRepository: implementasi nyata dari UserRepository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository: inisialisasi repo dengan dependency DB
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetByID(id int) (model.User, error) {
	var u model.User
	err := r.db.QueryRow("SELECT ID, Fullname, Email FROM Users WHERE ID = ?", id).
		Scan(&u.ID, &u.Fullname, &u.Email)
	return u, err
}

func (r *userRepository) GetByEmail(email string) (model.User, error) {
	var u model.User
	err := r.db.QueryRow("SELECT ID, Fullname, Email, Password FROM Users WHERE Email = ?", email).
		Scan(&u.ID, &u.Fullname, &u.Email, &u.Password)
	return u, err
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	result, err := r.db.Exec("INSERT INTO Users (Fullname, Email, Password) VALUES (?, ?, ?)",
		user.Fullname, user.Email, user.Password)
	if err != nil {
		return model.User{}, err
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)
	return user, nil
}

func (r *userRepository) Update(id int, user model.User) (model.User, error) {
	_, err := r.db.Exec("UPDATE Users SET Fullname = ?, Email = ? WHERE ID = ?",
		user.Fullname, user.Email, id)
	if err != nil {
		return model.User{}, err
	}
	user.ID = id
	return user, nil
}

func (r *userRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM Users WHERE ID = ?", id)
	return err
}
