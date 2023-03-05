package common

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"gorm.io/gorm"
)

func GetMainDb(sc goservice.ServiceContext) *gorm.DB {
	db := sc.MustGet(DBMain).(*gorm.DB)

	return db
}
