package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-gin-crud/domain"

	_courseHttpDelivery "go-gin-crud/course/delivery/http"
	_courseHttpDeliveryMiddleware "go-gin-crud/course/delivery/http/middleware"
	_courseRepo "go-gin-crud/course/repository/mysql"
	_courseUcase "go-gin-crud/course/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`release`) {
		gin.SetMode(gin.ReleaseMode)
		log.Println("Service RUN on RELEASE mode")
	} else {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("charset", "utf8mb4")
	val.Add("parseTime", "1")
	val.Add("loc", "Local")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := dbConn.DB()
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	dbConn.AutoMigrate(&domain.Course{}, &domain.CourseContent{})

	g := gin.New()
	middL := _courseHttpDeliveryMiddleware.InitMiddleware()
	g.Use(middL.CORS())

	courseRepo := _courseRepo.NewMysqlCourseRepository(dbConn)

	courseUC := _courseUcase.NewCourseUsecase(courseRepo)
	_courseHttpDelivery.NewCourseHandler(g, courseUC)

	log.Fatal(g.Run(viper.GetString("server.address")))
}
