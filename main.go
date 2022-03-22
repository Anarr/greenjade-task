package main

import (
	"github.com/Anarr/greenjade/config"
	"github.com/Anarr/greenjade/db"
	"github.com/Anarr/greenjade/route"
	"log"
	"net/http"
)

const defaultEnv = "development"

func main() {

	conf, err := config.Load(getEnv())

	must(err)

	conn, err := db.NewConnection(&db.AwsConnection{
		ApiKey:    conf.GetString("aws.access_key"),
		SecretKey: conf.GetString("aws.secret_key"),
		Region:    conf.GetString("aws.region"),
	})

	must(err)

	router := route.NewRouter(conn)
	//TOdo create new method like app.Run(mux) mux = router
	log.Fatal(http.ListenAndServe(":5001", router))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//return defult env for application
//can handle env from cli fo future
func getEnv() string {
	return defaultEnv
}
