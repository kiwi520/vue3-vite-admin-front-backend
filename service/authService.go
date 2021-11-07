package service

import (
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/respository"
	"log"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository respository.UserRepository
}

func (a authService) VerifyCredential(email string, password string) interface{} {
	res := a.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparePassword := CheckPasswordHash(password, v.Password)

		if v.Email == email && comparePassword {
			return res
		} else {
			return false
		}
	}
	return false
}

func (a authService) CreateUser(user dto.RegisterDTO) entity.User {

	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := a.userRepository.InsertUser(userToCreate)

	return res
}

func (a authService) FindByEmail(email string) entity.User {
	return a.userRepository.FindByEmail(email)
}

func (a authService) IsDuplicateEmail(email string) bool {
	res := a.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func NewAuthService(userRep respository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
