package Handler

import (
	"com.vic/DataAccess"
	"com.vic/DataEntity"
)

func GetUserInfo(userId string) *DataEntity.Tbl_User {
	user := DataAccess.GetUserInfo(userId)
	return user
}

//func GetPosts(userId string) *[]DataEntity.UserPosts {
//	rows:=DataAccess.GetPosts(userId)
//	return rows
//}
