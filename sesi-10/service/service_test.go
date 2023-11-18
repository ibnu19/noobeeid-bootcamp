package service

import (
	"database/sql"
	"sesi-10/database"
	"sesi-10/repository"
	"testing"
)

var dbSql *sql.DB
var svc service

func init() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	dbSql = db
	repo := repository.NewRepository(db)
	svc = NewService(repo)
}

func TestService(t *testing.T) {
	type testcase struct {
		title              string
		expectedErr        error
		expectedLenProduct int
		before             func()
	}

	testcases := []testcase{
		{
			title:              "success",
			expectedErr:        nil,
			expectedLenProduct: 1,
			before: func() {
				dbSql.Exec(`DELETE FROM products`)

				dbSql.Exec(`
				INSERT INTO products(id,name,category,price,stock)
				VALUES(1,'Nasi Goreng','Food',21000,50)
				`)
			},
		},
		{
			title:              "error not found",
			expectedErr:        ErrNotFound,
			expectedLenProduct: 0,
			before: func() {
				dbSql.Exec(`DELETE FROM products`)
			},
		},
	}

	for _, test := range testcases {
		t.Run(test.title, func(t *testing.T) {
			test.before()

			product, err := svc.GetAll()
			if err != test.expectedErr {
				t.Errorf("expected:%v, actual:%v", test.expectedErr, err)
			}

			if product != nil {
				if len(product) != test.expectedLenProduct {
					t.Errorf("expected:%v, actual:%v", test.expectedLenProduct, len(product))
				}
			}
		})
	}
}
