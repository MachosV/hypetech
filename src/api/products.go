package api

import (
	"api/apiutils"
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
	direction := r.FormValue("direction")
	pivot := r.FormValue("pivot")
	offsetBegin := r.FormValue("offset_begin")
	offsetEnd := r.FormValue("offset_end")
	query := apiutils.BuildQuery(direction, pivot, offsetBegin, offsetEnd)
	db := data.GetDbHandler()
	res, err := db.Query(query)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error")
		return
	}
	var productArray []models.Product
	var product models.Product
	id := -1
	maxid := id
	minid := id
	if res.Next() {
		res.Scan(
			&id,
			&product.Serial,
			&product.Name,
			&product.Description,
			&product.Quantity,
		)
		productArray = append(productArray, product)
		minid = id
		maxid = id
	}
	for res.Next() {
		res.Scan(
			&id,
			&product.Serial,
			&product.Name,
			&product.Description,
			&product.Quantity,
		)
		productArray = append(productArray, product)
		//set metadata for unsorted paging
		if maxid < id {
			maxid = id
		}
		if minid > id {
			minid = id
		}
	}
	var data Context
	data.Products = productArray
	data.Metadata = make(map[string]interface{})
	data.Metadata["maxid"] = maxid
	data.Metadata["minid"] = minid
	JSONdata, err := json.Marshal(data)
	fmt.Fprintf(w, "%s", JSONdata)
}

func notHandled(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method not yet implemented")
}
