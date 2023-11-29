package config

import (
	"fmt"
	"os"

	"orders/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	godotenv.Load()

	var (
		host     = os.Getenv("PGHOST")
		port     = os.Getenv("PGPORT")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	psqlInfo := fmt.Sprintf(
		`host=%s 
		port=%s 
		user=%s
		password=%s 
		dbname=%s 
		sslmode=disable`,
		host, port, user, password, dbname,
	)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to database")
	db.Debug().AutoMigrate(&model.Orders{}, &model.Items{})
}

func GetDB() *gorm.DB {
	return db
}
