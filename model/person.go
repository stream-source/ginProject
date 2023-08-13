package model

// Person 定义与数据库映射的模型
type Person struct {
	UserId   int64 `gorm:"column:user_id" gorm:"primaryKey"`
	username string
	sex      string
	email    string
}
