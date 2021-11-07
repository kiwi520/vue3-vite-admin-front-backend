package respository

import (
	"golang.org/x/crypto/bcrypt"
	"golang_api/entity"
	"gorm.io/gorm"
	"log"
)



type UserRepository interface {
	InsertUser( user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
    IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
	UserList() []entity.User
}

type userConnection struct {
	connection *gorm.DB

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
