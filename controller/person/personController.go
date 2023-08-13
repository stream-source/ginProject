package person

import (
	"ginProject/model"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func CreatePerson(c *gin.Context) {

}

func insertPerson(person *model.Person, db *gorm.DB) {
	db.Create(person)
}
