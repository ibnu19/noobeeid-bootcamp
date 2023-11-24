package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	APP_PORT    = os.Getenv("APP_PORT")
	DB          *sql.DB
)

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id := r.URL.Query().Get("id")

	var p Product

	query := `
		SELECT id,name,category,price,stock 
		FROM products
		WHERE id=$1
	`
	err := DB.QueryRow(query, id).Scan(&p.Id, &p.Name, &p.Category, &p.Price, &p.Stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		query := "SELECT id,name,category,price,stock FROM products ORDER BY id"
		rows, err := DB.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer rows.Close()

		var products []Product
		for rows.Next() {
			var p Product
			err := rows.Scan(&p.Id, &p.Name, &p.Category, &p.Price, &p.Stock)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			products = append(products, p)
		}

		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		name := r.FormValue("name")
		category := r.FormValue("category")
		price := r.FormValue("price")
		stock := r.FormValue("stock")

		query := `
			INSERT INTO products(name,category,price,stock)
			VALUES($1,$2,$3,$4)
		`

		_, err := DB.Exec(query, name, category, price, stock)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte("create product success"))
	}
}

func main() {
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db
	log.Println("database connected...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", GetProductById)
	mux.HandleFunc("/products", GetAllProducts)
	mux.HandleFunc("/products/create", CreateProduct)

	if APP_PORT == "" {
		APP_PORT = "3000"
	}

	http.ListenAndServe(":"+APP_PORT, mux)
}
