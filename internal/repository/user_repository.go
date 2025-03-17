package repository

import (
	"errors"
	"slices"
	"whois-client/internal/models"
)

type UserRepository struct {
	users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []models.User{
			{ID: 1, Contacts: []models.Contact{
				{ID: 1, Type: models.ContactTypeFullName, Value: "Vasya"},
				{ID: 2, Type: models.ContactTypeInstagramNick, Value: "@Vasya"},
				{ID: 3, Type: models.ContactTypePhone, Value: "+375291231212"},
			}},
		},
	}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	return r.users, nil
}

func (r *UserRepository) GetByID(id uint64) (*models.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return &u, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepository) Create(user *models.User) error {
	r.users = append(r.users, *user)
	return nil
}

func (r *UserRepository) Update(id uint64, user *models.User) error {
	for i, u := range r.users {
		if u.ID == id {
			r.users[i] = *user
			return nil
		}
	}

	return errors.New("user not found")
}

func (r *UserRepository) Delete(id uint64) error {
	for i, u := range r.users {
		if u.ID == id {
			r.users = slices.Delete(r.users, i, i+1)
			return nil
		}
	}

	return errors.New("user not found")
}
