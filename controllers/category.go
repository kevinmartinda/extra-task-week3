package controllers

import (
	"net/http"
	"log"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kevinmartinda/extra-task-week3/models"
	"github.com/kevinmartinda/extra-task-week3/config"
)

func FindAllCategories(res http.ResponseWriter, req *http.Request) {
	log.Printf("/GET categories")
	var category   	models.Category
	var categories  []models.Category
	var response 	models.CategoryResponse

	db := config.Connect()
	defer db.Close()

	records, err := db.Query("SELECT * FROM tb_kategori")
	
	if err != nil {
		log.Print(err)
	}

	for records.Next() {
		if err := records.Scan(&category.Id, &category.Nama); err != nil {
			log.Fatal(err.Error())
		} else {
			categories = append(categories, category)
		}
	}

	response.Status = 200
	response.Message = "success retrieve data"
	response.Data = categories

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(response)
}

func FindCategoryById (res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var category models.Category
	var categories []models.Category
	var response models.CategoryResponse

	db := config.Connect()
	defer db.Close()

	records, err := db.Query("SELECT * FROM tb_kategori WHERE id = ?", id)

	if err != nil {
		log.Print(err)
	}
	
	for records.Next() {
		if err := records.Scan(&category.Id, &category.Nama); err != nil {
			log.Fatal(err.Error())
		} else {
			categories = append(categories, category)
		}
	}
	response.Status = 200
	response.Message = "success retrieve data"
	response.Data = categories

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(response)
}

func CreateCategory(res http.ResponseWriter, req *http.Request) {
	db := config.Connect()
	defer db.Close()

	var category models.CategoryBody
	err := json.NewDecoder(req.Body).Decode(&category)

	log.Printf("finding body..")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("executing..")


	records, err := db.Exec("INSERT INTO tb_kategori (nama) VALUES (?)", category.Nama)

	if err != nil {
		log.Fatal(err.Error())
	}

	id, _ := records.LastInsertId()

	str := strconv.FormatInt(id, 10)

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(map[string]string{"status":"200", "message":"success add category","data":str})	
}

func FindCategoryByIdAndUpdate(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	stringId := vars["id"]

	db := config.Connect()
	defer db.Close()

	id, _ := strconv.ParseInt(stringId, 10, 64)

	var category models.CategoryBody
	err := json.NewDecoder(req.Body).Decode(&category)

	log.Printf("finding body..")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("executing..update")
	log.Printf(category.Nama)

	records, err := db.Exec("UPDATE tb_kategori SET nama = ? WHERE id = ?", category.Nama, id)

	if err != nil {
		log.Fatal(err.Error())
	}

	resid, _ := records.LastInsertId()

	str := strconv.FormatInt(resid, 10)

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(map[string]string{"status":"200", "message":"success edit category","data":str})	
}

func FindCategoryByIdAndDelete(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	db := config.Connect()
	defer db.Close()

	records, err := db.Exec("DELETE FROM tb_kategori WHERE id = ?", id)

	if err != nil {
		log.Fatal(err.Error())
	}

	resid, _ := records.LastInsertId()
	str := strconv.FormatInt(resid, 10)

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(map[string]string{"status":"200", "message":"success delete category","data":str})
}