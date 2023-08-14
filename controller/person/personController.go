package person

import (
	"ginProject/database"
	"ginProject/model"
	"github.com/gin-gonic/gin"
	"log"
)

func Controller(engine *gin.Engine) {
	person := engine.Group("/person")
	{
		person.POST("/create", CreatePerson)

	}
}

func CreatePerson(c *gin.Context) {
	//Post 请求直接获取json数据，并绑定到结构体对象上
	/**
	person := new(model.Person)
	person := &model.Person{}
	以上两种写法均是&地址，创建新的地址，故方法调用则insertPerson(person)直接用变量名传递即可，无法再次&取地址
	因为上面已经为变量分配内存地址
	var person model.Person
	insertPerson(&person)
	这种写法属于正常操作
	*/
	//person := new(model.Person)
	//person := &model.Person{}
	var person model.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		log.Panicln("解析json格式错误")
	}
	log.Printf("入库数据为，姓名：%s, 邮箱：%s\n", person.Username, person.Email)
	log.Printf("原始数据内存地址：%p\n", &person)
	insertPerson(&person)

}

func insertPerson(person *model.Person) {
	log.Printf("传递一次数据内存地址：%d\n", &person)
	//tx := database.DB.Create(&person)
	if err := database.DB.Table("person").Create(person).Error; err != nil {
		log.Println(err)
	}
	log.Printf("主键%d", person.UserId)
}
