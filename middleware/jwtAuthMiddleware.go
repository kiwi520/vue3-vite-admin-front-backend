package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang_api/helper"
	"golang_api/service"
	"log"
	"net/http"
)

//validate token middleware
func AuthorizeJwt(service service.JwtService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			res := helper.BuildErrorResponse(
				"缺少Authorization参数",
				errors.New("not found token"),
				)

			context.AbortWithStatusJSON(http.StatusBadRequest,res)
			return
		}

		token,err := service.ValidateToken(authHeader)

		if err != nil {
			res := helper.BuildErrorResponse(
				"缺少Authorization参数",
				err,
			)
			context.AbortWithStatusJSON(http.StatusUnauthorized,res)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("claim[user_id]:", claims["user_id"])
			log.Println("claim[issuer]:", claims["issuer"])
		}else if ve, ok := err.(*jwt.ValidationError); ok { //官方写法招抄就行
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("错误的token")
				res := helper.BuildErrorResponse("错误的token",err)
				context.AbortWithStatusJSON(http.StatusUnauthorized,res)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println("token过期或未启用")
				res := helper.BuildErrorResponse("token过期或未启用",err)
				context.AbortWithStatusJSON(http.StatusUnauthorized,res)
			} else {
				fmt.Println("无法处理这个token", err)
				res := helper.BuildErrorResponse("无法处理这个token",err)
				context.AbortWithStatusJSON(http.StatusUnauthorized,res)
			}

		}
	}
}
