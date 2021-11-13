package repository

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/helper"
	"gorm.io/gorm"
	"log"
	"strings"
)



type UserRepository interface {
	InsertUser( user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
    IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
	UserList() []entity.User
	GetUserPermission(userID int64) dto.UserPermissionResponse
	GetUserButtonList(userID int64) []string
}

type userConnection struct {
	connection *gorm.DB

}

func (u userConnection) GetUserButtonList(userID int64) []string {
	var button string
	err := u.connection.Table("users as u").Select("r.permission").Joins("left join roles as r on u.role_id = r.id ").Joins("left join roles as r on u.role_id = r.id ").Where("u.id",userID).Find(&button).Error
	if err != nil{
		panic(err)
	}

	buttonList := strings.Split(button,",")

	fmt.Println("buttonList")
	fmt.Println(buttonList)
	fmt.Println("buttonList")

	return nil
}

func (u userConnection) GetUserPermission(userID int64) (data dto.UserPermissionResponse) {
	var userPermission string
	//err := u.connection.Table("users as u").Select("r.menu_json,r.button_string").Joins("left join roles as r on u.role_id = r.id ").Where("u.id",userID).Find(&userPermission).Error
	err := u.connection.Table("users as u").Select("r.permission").Joins("left join roles as r on u.role_id = r.id ").Where("u.id",userID).Find(&userPermission).Error
	if err != nil{
		panic(err)
	}

	var permissionList = []entity.Menu{}

	MenuJsonSlice:=strings.Split(userPermission,",")

	u.connection.Where("id in ?",MenuJsonSlice).Find(&permissionList)

	var buttonList = []string{}
	var menuList = []entity.Menu{}

	if len(permissionList) > 0 {
		for _,item := range permissionList{
			if item.Type == 1 {
				menuList = append(menuList,item)
			}else if item.Type == 2 {
				buttonList = append(buttonList,item.Code)
			}
		}
	}

	menuTreeList:= helper.GetMenuTree(menuList,0)

	data.MenuTreeList = menuTreeList
	data.ButtonString = strings.Join(buttonList, ",")

	return data
}

func (u userConnection) UserList() []entity.User {
	var userList = []entity.User{}
	u.connection.Find(&userList)
	return userList
}

func (u userConnection) InsertUser(user entity.User) entity.User {
	password, err := HashPassword([]byte(user.Password))
	if err != nil {
		log.Println(err)
		panic("加密密码失败")
	}
	user.Password = password

	u.connection.Save(&user)
	return user
}

func (u userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != " " {
		password, err := HashPassword([]byte(user.Password))
		if err != nil {
			log.Println(err)
			panic("加密密码失败")
		}
		user.Password = password
	}else {
		var tmpUser entity.User
		u.connection.Find(&tmpUser,user.ID)
		user.Password = tmpUser.Password
	}


	//u.connection.Save(&user)

	u.connection.Save(&user)
	return user
}

func (u userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := u.connection.Where("email = ? ",email).Take(&user)

	if res.Error == nil {
		return user
	}else {
		return res.Error
	}
}

func (u userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return  u.connection.Where("email = ? ",email).Take(&user)

}

func (u userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	res:= u.connection.Where("email = ? ",email).Take(&user)

	if res.Error == nil {
		return user
	}else {
		return entity.User{}
	}
}

func (u userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	res:= u.connection.Preload("Books").Find(&user,userID)

	if res.Error == nil {
		return user
	}else {
		return entity.User{}
	}
}


func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func HashPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
