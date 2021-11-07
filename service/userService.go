package service

import (
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/respository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	List() []entity.User
}

type userService struct {
	userRepository respository.UserRepository
}

func (u userService) List() []entity.User {
	return u.userRepository.UserList()
}

func (u userService) Update(user dto.UserUpdateDTO) entity.User {
	userToCreate := entity.User{}
	//err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	//if err != nil {
	//	log.Fatalf("Failed map %v", err)
	//}

	userToCreate.ID= uint(user.ID)
	userToCreate.Name= user.Name
	userToCreate.Email= user.Email
	userToCreate.Password= user.Password


	println("userToCreate")
	println(userToCreate.ID)
	println(userToCreate.Name)
	println(userToCreate.Email)
	println("userToCreate")
	updateUser := u.userRepository.UpdateUser(userToCreate)

	return  updateUser
}

func (u userService) Profile(userID string) entity.User {
	user := u.userRepository.ProfileUser(userID)

	return user
}

func NewUserService(userRepository respository.UserRepository) UserService {
	return  &userService{
		userRepository,
	}
}