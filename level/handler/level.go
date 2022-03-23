package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Anarr/greenjade/internal/response"
	"github.com/Anarr/greenjade/level/repository"
	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/julienschmidt/httprouter"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Item struct {
	Id     string  `json:"id"`
	Levels [][]int `json:"levels"`
}

//StoreLevelRequest validate store new level
type StoreLevelRequest struct {
	Level [][]int `valid:"required" json:"level"`
}

//UpdateLevelRequest validate update new level
type UpdateLevelRequest struct {
	Id    string  `json:"id"`
	Level [][]int `valid:"required" json:"level"`
}

//Validate validate StoreLevelRequest
//check dimension max len
//check dimensions are quadratic
//check each dimension levels value 0, 1 or 2
func (slr *StoreLevelRequest) Validate() error {

	if len(slr.Level) == 0 {
		return errors.New("level can not be empty")
	}

	maxLen := 100
	dimensionLen := len(slr.Level[0])

	for _, v := range slr.Level {
		//check each dimension size less or equal 100
		if len(v) > maxLen {
			return errors.New(fmt.Sprintf("dimension size must be less %d", maxLen))
		}

		//check each dimension len is equal
		if len(v) != dimensionLen {
			return errors.New("each dimension must the same len")
		}

		dimensionLen = len(v)

		//check current dimension elements.

		for _, e := range v {
			if e > 2 || e < 0 {
				return errors.New("dimension elements value must be 0, 1 and 2")
			}
		}
	}

	return nil
}

//Convert slice of int to slice of string
func (slr *StoreLevelRequest) toDynamoddbList() []*dynamodb.AttributeValue {
	var xs []*dynamodb.AttributeValue

	for _, v := range slr.Level {
		var l dynamodb.AttributeValue
		var xss []*dynamodb.AttributeValue

		for _, e := range v {
			str := strconv.Itoa(e)
			value := &dynamodb.AttributeValue{
				N: aws.String(str),
			}
			xss = append(xss, value)
		}

		l.L = xss
		xs = append(xs, &l)
	}

	return xs
}

//GetRandID generate random id for new level
func (slr *StoreLevelRequest) GetRandID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000)
}

//GetByIdHandler handle retrieve single level request
func GetByIdHandler(db *dynamodb.DynamoDB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		//check id must be integer
		if !govalidator.IsNumeric(params.ByName("id")) {
			response.Error(w, "opps")
			return
		}

		//retrieve data from AWS
		level, err := repository.GetById(db, params.ByName("id"))

		if err != nil {
			response.Error(w, "Not exists level")
			return
		}

		response.Success(w, level)
		return
	}
}

//StoreHandler handle store new level request
func StoreHandler(db *dynamodb.DynamoDB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var res StoreLevelRequest
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&res); err != nil {
			fmt.Println(err)
			response.Error(w, "Could not handle user request.")
			return
		}

		if err := res.Validate(); err != nil {
			response.Error(w, err.Error())
			return
		}

		id, err := repository.Store(db, res.GetRandID(), res.Level)

		if err != nil {
			response.Error(w, err.Error())
			return
		}

		result := make(map[string]int)
		result["id"] = id

		response.Success(w, result)
		return
	}
}

//UpdateHandler handle update level request
func UpdateHandler(db *dynamodb.DynamoDB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var res StoreLevelRequest
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&res); err != nil {
			fmt.Println(err)
			response.Error(w, "Could not handle user request.")
			return
		}

		if err := res.Validate(); err != nil {
			response.Error(w, err.Error())
			return
		}

		if err := repository.Update(db, params.ByName("id"), res.Level); err != nil {
			response.Error(w, err.Error())
			return
		}

		response.Success(w, "successfully updated")
		return
	}
}

//RemoveHandler handle remove level request
func RemoveHandler(db *dynamodb.DynamoDB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		err := repository.Remove(db, params.ByName("id"))

		if err != nil {
			response.Error(w, err.Error())
			return
		}

		response.Success(w, "successfully deleted")
		return
	}
}
