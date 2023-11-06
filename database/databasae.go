package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DBconn 변수는 데이터베이스 연결을 저장합니다.
var (
	DBconn *gorm.DB
)
