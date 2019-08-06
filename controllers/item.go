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

func FindAllItem(res http.ResponseWriter, req *http.Request) {
	log.Printf("/GET items")
	var item   	models.Item
	var items  []models.Item
	var response 	models.ItemResponse

	db := config.Connect()
	defer db.Close()

	records, err := db.Query("SELECT tb_item.id, tb_item.item_id, tb_kategori.nama FROM tb_item JOIN tb_kategori ON tb_item.kategori_id = tb_kategori.id")
		if err != nil {
			log.Print(err)
		}

		for records.Next() {
			if err := records.Scan(&item.Id, &item.ItemId, &item.Nama); err != nil {
				log.Fatal(err.Error())
			} else {
				items = append(items, item)
			}
		}

	response.Status = 200
	response.Message = "success retrieve data"
	response.Data = items

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(response)
}

func FindItemById (res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var item models.Item
	var items []models.Item
	var response models.ItemResponse

	db := config.Connect()
	defer db.Close()

	records, err := db.Query("SELECT tb_item.id, tb_item.item_id, tb_kategori.nama FROM tb_item JOIN tb_kategori ON tb_item.kategori_id = tb_kategori.id WHERE tb_item.id = ?", id)

	if err != nil {
		log.Print(err)
	}
	
	for records.Next() {
		if err := records.Scan(&item.Id, &item.ItemId, &item.Nama); err != nil {
			log.Fatal(err.Error())
		} else {
			items = append(items, item)
		}
	}
	response.Status = 200
	response.Message = "success retrieve data"
	response.Data = items

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(response)
}

func CreateItem(res http.ResponseWriter, req *http.Request) {
	db := config.Connect()
	defer db.Close()

	var item models.ItemBody
	err := json.NewDecoder(req.Body).Decode(&item)

	log.Printf("finding body..")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("executing..")


	records, err := db.Exec("INSERT INTO tb_item (item_id, kategori_id) VALUES (?, ?)", item.ItemId, item.CategoryId)

	if err != nil {
		json.NewEncoder(res).Encode(map[string]string{"status":"200", "message":"It seems that item id you entered already exist"})
	} else {
	
		id, _ := records.LastInsertId()
	
		str := strconv.FormatInt(id, 10)
	
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Content-Type", "application/json")
	
		json.NewEncoder(res).Encode(map[string]string{"status":"200", "message":"success add category","data":str})
	}	
}

func FindItemByIdAndUpdate(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	stringId := vars["id"]

	db := config.Connect()
	defer db.Close()

	id, _ := strconv.ParseInt(stringId, 10, 64)

	var item models.ItemBody
	err := json.NewDecoder(req.Body).Decode(&item)

	log.Printf("finding body..")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("executing..update")

	records, err := db.Exec("UPDATE tb_item SET item_id = ?, kategori_id = ? WHERE id = ?", item.ItemId, item.CategoryId, id)

	if err != nil {
		log.Fatal(err.Error())
	}

	resid, _ := records.LastInsertId()

	str := strconv.FormatInt(resid, 10)

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(map[string]string{"status":"200", "message":"success edit item","data":str})	
}

func FindItemByIdAndDelete(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	db := config.Connect()
	defer db.Close()

	records, err := db.Exec("DELETE FROM tb_item WHERE id = ?", id)

	if err != nil {
		log.Fatal(err.Error())
	}

	resid, _ := records.LastInsertId()
	str := strconv.FormatInt(resid, 10)

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(map[string]string{"status":"200", "message":"success delete category","data":str})
}