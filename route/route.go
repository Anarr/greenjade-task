package route

import (
	"github.com/Anarr/greenjade/level/handler"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/julienschmidt/httprouter"
)

//NewRouter create new httprouter.Router instance
//register app all routes here
func NewRouter(db *dynamodb.DynamoDB) *httprouter.Router {
	router := httprouter.New()

	router.GET("/levels/:id", handler.GetByIdHandler(db))
	router.POST("/levels", handler.StoreHandler(db))
	router.PATCH("/levels/:id", handler.UpdateHandler(db))
	router.DELETE("/levels/:id", handler.RemoveHandler(db))
	return router
}
