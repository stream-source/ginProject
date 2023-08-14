package model

// Person 定义与数据库映射的模型
type Person struct {
	UserId   int64  `gorm:"column:user_id;type:AUTO_INCREMENT" gorm:"primaryKey" json:"userId" form:"userId"`
	Username string `gorm:"column:username" json:"username" form:"username"`
	Sex      string `gorm:"column:sex" json:"sex" form:"sex"`
	Email    string `gorm:"column:email" json:"email"  form:"email"`
}

// 重写方法，定义表名
func (Person) TableName() string {
	return "person"
}
