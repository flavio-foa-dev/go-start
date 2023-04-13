package models

import (
	"app/myapp/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProductsAll() []Product {
	db := db.ConnectionDB()
	defer db.Close()

	getProductsAll, err := db.Query("SELECT * FROM produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}
	for getProductsAll.Next() {
		var id int
		var name string
		var description string
		var price float64
		var quantity int

		err := getProductsAll.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()

	return products
}

func InsertProduct(name string, description string, price float64, quantity int) {
	db := db.ConnectionDB()
	defer db.Close()

	insertProduct, err := db.Prepare("INSERT INTO produtos (name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, quantity)
}

func DeleteProduct(id string) {
	db := db.ConnectionDB()
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
}

func EditProduct(id string) Product {
	db := db.ConnectionDB()
	defer db.Close()

	productById, err := db.Query("SELECT * FROM produtos WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for productById.Next() {
		var id, quantity int
		var price float64
		var name, description string

		err = productById.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	return product
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := db.ConnectionDB()
	defer db.Close()

	prepare, err := db.Prepare("update produtos set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	prepare.Exec(name, description, price, quantity, id)
}
