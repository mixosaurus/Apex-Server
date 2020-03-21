package DataEntity

type Tbl_User struct {
	UserId   uint   `gorm:"column:userid;primary_key"`
	UserName string `gorm:"column:username"`
}

//// 设置User的表名
//func (Tbl_User) TableName() string {
//	return "tbl_users"
//}
