package serves

import (
	"context"
	"fmt"

	"github.com/getfarook/go_grpc_crud_server/prodpb"
	"github.com/getfarook/go_grpc_crud_server/salespro/dbcon"
)

type Prodserver struct{}

func (*Prodserver) AddProduct(ctx context.Context, req *prodpb.Product) (*prodpb.ProdId, error) {
	fmt.Printf("Add Product function was invoked with : %v \n", req)
	name := req.GetName()
	price := req.GetPrice()
	details := req.GetDetails()

	db := dbcon.DbHandle()
	defer db.Close()

	sqlStatement := "INSERT INTO products (name, price, details) VALUES ($1, $2, $3) RETURNING id"
	var id int32 = 0

	err := db.QueryRow(sqlStatement, name, price, details).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

	resp := &prodpb.ProdId{

		Id: id,
	}

	return resp, nil
}

func (*Prodserver) DeleteProduct(ctx context.Context, req *prodpb.ProdId) (*prodpb.DelProdResponse, error) {
	id := req.GetId()

	db := dbcon.DbHandle()
	defer db.Close()

	sqlStatement := "DELETE FROM products WHERE id = $1;"
	res, err1 := db.Exec(sqlStatement, id)
	if err1 != nil {
		panic(err1)
	}

	count, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}

	fmt.Print("\n No. of Rows deleted were: ")
	fmt.Println(count)

	resp := &prodpb.DelProdResponse{} //empty dummy response
	return resp, nil
}

func (*Prodserver) UpdateProduct(ctx context.Context, req *prodpb.Product) (*prodpb.DelProdResponse, error) {
	id := req.GetId()
	name := req.GetName()
	price := req.GetPrice()
	details := req.GetDetails()

	db := dbcon.DbHandle()
	defer db.Close()

	sqlStatement := "UPDATE products SET name = $2, price = $3, details = $4 WHERE id = $1;"

	res, err1 := db.Exec(sqlStatement, id, name, price, details)
	if err1 != nil {
		panic(err1)
	}

	count, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}

	fmt.Print("\n No of rows updated were: ")
	fmt.Println(count)

	resp := &prodpb.DelProdResponse{} //empty dummy response
	return resp, nil
}

func (*Prodserver) GetProducts(ctx context.Context, req *prodpb.GetProdRequest) (*prodpb.Products, error) {
	fmt.Printf("Get Products function is invoked")

	var db = dbcon.DbHandle()
	defer db.Close()

	rows, err1 := db.Query("SELECT id, name, price, details FROM products ")
	if err1 != nil {
		// handle this error better than this
		panic(err1)
	}
	defer rows.Close()

	loadedProducts := []*prodpb.Product{}
	for rows.Next() {
		var id int32
		var name string
		var price float64
		var details string
		err2 := rows.Scan(&id, &name, &price, &details)
		if err2 != nil {
			// handle this error
			panic(err2)
		}

		newProduct := &prodpb.Product{
			Id:      id,
			Name:    name,
			Price:   price,
			Details: details,
		}

		loadedProducts = append(loadedProducts, newProduct)

	}

	products := &prodpb.Products{
		MyProducts: loadedProducts,
	}

	err3 := rows.Err()
	if err3 != nil {
		panic(err3)
	}

	return products, nil

}
