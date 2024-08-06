package bootstrap

import (
	"arquitectura/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func newDatabase() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	hostDb := os.Getenv("DB_HOST")
	portDb := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_DB")
	//sslMode := os.Getenv("DB_SSLMODE")

	/*config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
	hostDb, portDb, user, password, dbname)*/

	dsn := "host=" + hostDb + " port=" + portDb + " user=" + user + " password=" + password + " dbname=" + dbname
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&model.User{})

	/*db, err := sql.Open("pgx", config)
	if err != nil {
		log.Fatalf("could not create a connection to the database, err: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}*/

	return db
}
