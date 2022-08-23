package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

func main() {
	engine := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN())
	if err != nil {
		fmt.Printf("driver register error.")
	}
	defer cleanup()

	db, err := sql.Open(
		"cloudsql-postgres",
		"host={} user={} password={} dbname={} sslmode=disable",
	)
	if err != nil {
		fmt.Printf("create connection error.")
	}
	defer db.Close()

	engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello")
	})
	engine.GET("/a", func(context *gin.Context) {

		rows, err := db.Query("select 1")
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			return
		}
		defer rows.Close()

		var count string
		for rows.Next() {
			if err := rows.Scan(&count); err != nil {
				context.String(http.StatusBadRequest, err.Error())
				return
			}
		}
		context.String(http.StatusOK, "%#+v", count)
	})
	err = engine.Run("0.0.0.0:" + port())
	if err != nil {
		fmt.Errorf("error : %+v", err)
	}
}

func port() string {
	p := os.Getenv("PORT")
	fmt.Printf("port : %s\n", p)
	if p == "" {
		p = "8080"
	}
	fmt.Printf("port2 : %s\n", p)
	return p
}
