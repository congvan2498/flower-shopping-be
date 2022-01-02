package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"os"
	"togo/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()

	var dsn string
	//connectingString:= "sqlserver://username:password@host:port?database=nameDb"
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	//dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", dbHost, dbUsername, dbPassword, dbDatabase, dbPort)

	//Connect()

	dsn = fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)

	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Println(err)
		panic("Failed to connect database")
	}

	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Task{})
	database.AutoMigrate(&model.Product{})

	// auto migrate order

	database.Migrator().CreateConstraint(&model.Address{}, "address")
	database.Migrator().CreateConstraint(&model.Address{}, "pk_order_address")

	database.Migrator().CreateConstraint(&model.Department{}, "Users")
	database.Migrator().CreateConstraint(&model.Department{}, "pk_department_user")

	database.AutoMigrate(&model.Order{})
	database.AutoMigrate(&model.Address{})


	if err != nil {
		panic(err.Error())
	}

	DB = database
	return DB
}

var db *sql.DB
var server = "shopping-card.database.windows.net"
var port = 1433
var user = "shopping-card"
var password = "Van123456"
var database = "shopping-card1"

func Connect() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
}
