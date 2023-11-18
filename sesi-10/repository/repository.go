package repository

import (
	"database/sql"
	"sesi-10/model"
)

type ProductRepository interface {
	GetAll() (products []model.Product, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

func (repository repository) GetAll() (products []model.Product, err error) {
	query := `SELECT id,name,category,price,stock FROM products`
	rows, err := repository.db.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var p model.Product
		err = rows.Scan(&p.Id, &p.Name, &p.Category, &p.Price, &p.Stock)
		if err != nil {
			return
		}
		products = append(products, p)
	}

	return
}
