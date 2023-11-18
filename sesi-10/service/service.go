package service

import (
	"database/sql"
	"errors"
	"sesi-10/model"
	"sesi-10/repository"
)

var (
	ErrNotFound       = errors.New("data not found")
	ErrInternalServer = errors.New("internal server error")
)

type service struct {
	repo repository.ProductRepository
}

func NewService(repo repository.ProductRepository) service {
	return service{
		repo: repo,
	}
}

func (s service) GetAll() (products []model.Product, err error) {
	products, err = s.repo.GetAll()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		} else {
			return nil, ErrInternalServer
		}
	}

	if len(products) == 0 {
		return nil, ErrNotFound
	}
	return
}
