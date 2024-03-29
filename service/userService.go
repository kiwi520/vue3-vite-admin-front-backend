package service

import (
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	List() []entity.User
	GetUserPermission(userID int64) dto.UserPermissionResponse
	GetUserButtonList(userID int64) []string
}

type userService struct {
	userRepository repository.UserRepository
}

func (u userService) GetUserButtonList(userID int64) []string {
	return u.userRepository.GetUserButtonList(userID)
}

func (u userService) GetUserPermission(userID int64) dto.UserPermissionResponse {
	return u.userRepository.GetUserPermission(userID)
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

func NewUserService(userRepository repository.UserRepository) UserService {
	return  &userService{
		userRepository,
	}
}