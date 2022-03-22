package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Anarr/greenjade/internal/response"
	"github.com/Anarr/greenjade/level/repository"
	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Item struct {
	Id string `json:"id"`
	Levels [][]int `json:"levels"`
}

//StoreLevelRequest validate store new level
type StoreLevelRequest struct {
	Level [][]int `valid:"required" json:"level"`
}

//GetByIdHandler retrieve single level
func GetByIdHandler(db *dynamodb.DynamoDB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		//check id must be integer
		if !govalidator.IsNumeric(params.ByName("id")) {
			response.Error(w, "opps")
			return
		}

		//retrieve data from AWS
		level, err := repository.GetById(db, params.ByName("id"))

		if err != nil || level.WithEmptyId(){
			response.Error(w, "Not exists level")
			return
		}

		response.Success(w, level)
		return
	}
}

//StoreHandler store new level
func StoreHandler(db *dynamodb.DynamoDB) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var res StoreLevelRequest
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&res); err != nil {
			fmt.Println(err)
			response.Error(w, "Could not handle user request.")
			return
		}

		_, err := govalidator.ValidateStruct(res)

		if err != nil {
			response.Error(w, err.Error())
			return
		}

		fmt.Fprint(w, "can create")
	}
}

func UpdateHandler(db *dynamodb.DynamoDB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(w, "cool")
	}
}

//RemoveHandler remove level with given id
func RemoveHandler(db *dynamodb.DynamoDB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(w, fmt.Sprintf("Retrieve level #%s", params.ByName("id")))
	}
}