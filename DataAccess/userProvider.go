package DataAccess

import (
	"com.vic/DBHelper"
	"com.vic/DataEntity"
)

func GetUserInfo(userId string) *DataEntity.Tbl_User {
	db := DBHelper.NewInstance()
	defer db.Close()
	var user DataEntity.Tbl_User
	rows := db.Where("userid=?", userId).Find(&user)

	//rows:=db.First(&user)

	if rows.Error != nil {
		//fmt.Println(rows.Error.Error())
		panic(rows.Error.Error())
	}

	return &user
}

//func GetPosts(userId string) *[]DataEntity.UserPosts{
//	db:=DBHelper.NewInstance()
//	defer db.Close()
//
//	var user DataEntity.Tbl_User
//	var count int64
//
//	db.Where("userId=?","aaa").Find(&user)
//	user.UserName="donneemay"
//
//	db.Save(&user)
//	var posts []DataEntity.UserPosts
//	rows:=db.Table("tbl_posts as tp").
//		Select("tp.postsId,tp.postsMaker,tu.username as postsMakerName,tp.postsContent").
//		Joins("left join tbl_user tu on tp.postsMaker=tu.userId").
//		Where("tu.userId=?").
//		Order("tp.postId desc").
//		Find(&posts).
//		Count(&count)
//
//	err := rows.Error
//	if err!=nil{
//		fmt.Println(err)
//	}
//
//	//_=rows
//}
//
//
