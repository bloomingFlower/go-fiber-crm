package main

import (
	"fmt"

	"github.com/bloomingFlower/go-fiber-crm/database"
	"github.com/bloomingFlower/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/", lead.GetLeads)
	app.Get("", lead.Getlead)
	app.Post(lead.NewLead)
	app.Delete(lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBconn, err = database.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBconn.Close()
}
