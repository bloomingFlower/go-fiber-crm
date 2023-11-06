package main

import (
	"fmt"

	"github.com/bloomingFlower/go-fiber-crm/database"
	"github.com/bloomingFlower/go-fiber-crm/lead"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// setupRoutes 함수는 웹 서버의 라우트를 설정합니다.
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/", lead.GetLeads)         // 모든 Lead를 가져오는 라우트를 설정합니다.
	app.Get("/api/v1/lead/:id", lead.GetLead)       // 특정 Lead를 가져오는 라우트를 설정합니다.
	app.Post("/api/v1/lead", lead.NewLead)          // 새로운 Lead를 생성하는 라우트를 설정합니다.
	app.Delete("/api/v1/lead/:id", lead.DeleteLead) // 특정 Lead를 삭제하는 라우트를 설정합니다.
}

// initDatabase 함수는 데이터베이스를 초기화합니다.
func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "leads.db") // SQLite 데이터베이스에 연결합니다.
	if err != nil {
		panic("failed to connect database") // 데이터베이스 연결에 실패하면 패닉을 발생시킵니다.
	}
	fmt.Println("Connection Opened to Database") // 데이터베이스 연결 성공 메시지를 출력합니다.
	database.DBconn.AutoMigrate(&lead.Lead{})    // Lead 테이블을 자동으로 마이그레이션합니다.
	fmt.Println("Database Migrated")             // 데이터베이스 마이그레이션 성공 메시지를 출력합니다.
}

// main 함수는 웹 서버를 실행합니다.
func main() {
	app := fiber.New()            // 새로운 Fiber 앱을 생성합니다.
	initDatabase()                // 데이터베이스를 초기화합니다.
	setupRoutes(app)              // 라우트를 설정합니다.
	app.Listen(":3000")           // 3000번 포트에서 웹 서버를 실행합니다.
	defer database.DBconn.Close() // main 함수가 끝나면 데이터베이스 연결을 닫습니다.
}
