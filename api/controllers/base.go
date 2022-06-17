package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/imsujan276/untold-story/api/middlewares"
	"github.com/imsujan276/untold-story/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

var errList = make(map[string]string)

func (server *Server) Init(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	// If you are using mysql, i added support for you here(dont forgot to edit the .env file)
	if DbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(DbDriver, DBURL)
		if err != nil {
			log.Fatal("Error connecting to mysql:", err)
		} else {
			log.Printf("We are connected to the %s database", DbDriver)
		}
	} else if DbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(DbDriver, DBURL)
		if err != nil {
			log.Fatal("Error connecting to postgres: ", err)
		} else {
			log.Printf("We are connected to the %s database", DbDriver)
		}
	} else {
		log.Println("Unknown Driver")
	}

	//database migration
	server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.ResetPassword{},
		&models.Like{},
		&models.Comment{},
	)

	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())
	server.Router.Static("/static", "./static")

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
