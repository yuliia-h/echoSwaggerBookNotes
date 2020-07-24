package main

import (
	"database/sql"
	"echoSwaggerBookNotes/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	db := initDB("storage.db")
	migrate(db)

	// Create a new instance of Echo
	e := echo.New()

	e.File("/", "frontend/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// Start as a web server
	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	//откроем файл или создадим его
	db, err := sql.Open("sqlite3", filepath)
	// проверяем ошибки и выходим при их наличии
	if err != nil {
		panic(err)
	}
	// если ошибок нет, но не можем подключиться к базе данных,
	// то так же выходим
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sqDb := `
		CREATE TABLE IF NOT EXISTS notes(
		   id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		   name VARCHAR NOT NULL
   	);
   	`

	_, err := db.Exec(sqDb)
	// выходим, если будут ошибки с SQL запросом выше
	if err != nil {
		panic(err)
	}
}