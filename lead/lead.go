package lead

import (
	"errors"
	"net/http"

	"github.com/bloomingFlower/go-fiber-crm/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Lead라는 이름의 구조체를 정의합니다.
type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int64  `json:"phone"`
}

// 모든 Lead를 가져오는 함수를 정의합니다.
func GetLeads(c *fiber.Ctx) error {
	db := database.DBconn // 데이터베이스 연결을 가져옵니다.
	var leads []Lead      // Lead의 슬라이스를 선언합니다.
	db.Find(&leads)       // 모든 Lead를 찾아서 leads에 저장합니다.
	c.JSON(leads)         // leads를 JSON 형태로 응답합니다.
	return nil
}

// 특정 Lead를 가져오는 함수를 정의합니다.
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")  // 요청에서 id 파라미터를 가져옵니다.
	db := database.DBconn // 데이터베이스 연결을 가져옵니다.
	var lead Lead         // Lead 구조체를 선언합니다.
	db.Find(&lead, id)    // id에 해당하는 Lead를 찾아서 lead에 저장합니다.
	c.JSON(lead)          // lead를 JSON 형태로 응답합니다.
	return nil
}

// NewLead 함수는 새로운 Lead를 생성하는 역할을 합니다.
func NewLead(c *fiber.Ctx) error {
	db := database.DBconn                      // 데이터베이스 연결을 가져옵니다.
	lead := new(Lead)                          // 새로운 Lead 구조체를 생성합니다.
	if err := c.BodyParser(lead); err != nil { // 요청 본문을 파싱하여 lead에 저장합니다.
		c.Status(http.StatusServiceUnavailable).Send([]byte(err.Error())) // 에러가 발생하면 에러 메시지와 함께 503 응답을 보냅니다.
		return errors.New("Error parsing JSON")
	}
	db.Create(&lead) // 데이터베이스에 새로운 Lead를 생성합니다.
	c.JSON(lead)     // 생성된 Lead를 JSON 형태로 응답합니다.
	return nil
}

// DeleteLead 함수는 특정 Lead를 삭제하는 역할을 합니다.
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")  // 요청에서 id 파라미터를 가져옵니다.
	db := database.DBconn // 데이터베이스 연결을 가져옵니다.

	var lead Lead        // Lead 구조체를 선언합니다.
	db.First(&lead, id)  // id에 해당하는 Lead를 찾아서 lead에 저장합니다.
	if lead.Name == "" { // Lead가 존재하지 않으면 에러 메시지와 함께 500 응답을 보냅니다.
		c.Status(http.StatusInternalServerError).Send([]byte("No lead found with given ID"))
		return errors.New("No lead found with given ID")
	}
	db.Delete(&lead)                            // 데이터베이스에서 Lead를 삭제합니다.
	c.Send([]byte("Lead Successfully deleted")) // 삭제 성공 메시지를 응답합니다.
	return nil
}
