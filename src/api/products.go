package api

import (
	"data"
	"encoding/json"
	"fmt"
	"log"
	"middleware"
	"models"
	"mux"
	"net/http"
)

type Context struct {
	Products []models.Product       `json:"products"`
	Metadata map[string]interface{} `json:"metadata"`
}

func init() {
	mux := mux.GetMux()
	mux.HandleFunc("/products", middleware.WithMiddleware(productHandler,
		middleware.Time(),
	))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		getProducts(w, r)
	default:
		notHandled(w, r)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	db := data.GetDbHandler()
	res, err := db.Query("SELECT id,pserial,pname,pdesc,quantity FROM products LIMIT 10;")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error")
		return
	}
	var productArray []models.Product
	var product models.Product
	maxid := 0
	id := 0
	for res.Next() {
		res.Scan(
			&id,
			&product.Serial,
			&product.Name,
			&product.Description,
			&product.Quantity,
		)
		productArray = append(productArray, product)
		if id > maxid {
			maxid = id
		}
	}
	var result Context
	result.Products = productArray
	result.Metadata = make(map[string]interface{})
	result.Metadata["lastid"] = maxid
	productsJSON, err := json.Marshal(result)
	fmt.Fprintf(w, "%s", productsJSON)
}

func notHandled(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Error, method not handled")
}
