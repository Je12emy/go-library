package app

import (
	"encoding/json"
	"fmt"
	"library/logger"
	"library/service"
	"log"
	"net/http"
	"os"

	"library/domain"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func checkEnvVariables() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	if serverAddress == "" || serverPort == "" {
		logger.Error("Missing enviroment variables, shutting down server.")
		log.Fatal("Missing enviroment variables")
	}
}

// Start Check for .env variables, and wire the whole application
func Start() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file with godotenv.")
		panic("Error loading .env file.")
	}

	checkEnvVariables()

	// Create a new gorilla multiplexer
	router := mux.NewRouter()

	// Wire up the app
	bookRepo := domain.NewBookRepositoryDB(GetDBClient())
	bh := BookHandler{service.NewBookService(bookRepo)}

	// db := GetDBClient()
	router.HandleFunc("/books", bh.NewBook).
		Methods(http.MethodPost).
		Name("CreateBook")

	router.HandleFunc("/books/{book_id:[0-9]+}", bh.FindBook).
		Methods(http.MethodGet).
		Name("FindSingleBook")

	router.HandleFunc("/books/{book_id:[0-9]+}", bh.UpdateBook).
		Methods(http.MethodPut).
		Name("UpdateBook")

	router.HandleFunc("/books/{book_id:[0-9]+}", bh.DeleteBook).
		Methods(http.MethodDelete).
		Name("DeleteBook")

	router.HandleFunc("/books", bh.FindAllBooks).
		Methods(http.MethodGet).
		Name("FindAllBooks")

	serverAddress := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	address := fmt.Sprintf("%v:%v", serverAddress, serverPort)
	logger.Info("Started server with address: http://" + address)
	log.Fatal(http.ListenAndServe(address, router))
}

// WriteResponse : Writes a json response with a given code and data
func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}

// GetDBClient Creates a connection with a production db.
func GetDBClient() *gorm.DB {
	// Conect to sqllite test db
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate
	db.AutoMigrate(&domain.Book{})

	return db
}

// GetTestDBClient Creates a connection with the test db.
func GetTestDBClient() *gorm.DB {
	// Conect to sqllite test db
	db, err := gorm.Open(sqlite.Open("../library_test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate
	db.Migrator().DropTable(&domain.Book{})
	db.AutoMigrate(&domain.Book{})

	return db
}
