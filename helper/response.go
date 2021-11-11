package helper

import (
	"golang_api/dto"
	"golang_api/entity"
)

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Errors interface{} `json:"errors"`
	Data interface{} `json:"data"`
}

type Empty struct {
	
}

func BuildResponse(status int, message string, data interface{})  Response {
	res := Response{
		Status: status,
		Message: message,
		Errors: nil,
		Data: data,
	}

	return res
}

func BuildErrorResponse(message string,error error) Response {
	res := Response{
		Message: message,
		Errors: error.Error(),
	}

	return res
}


func GetMenuTree(list []entity.Menu, pid uint) []dto.MenuTree {
	var MenuTree []dto.MenuTree
	for _,val := range list {
		if val.ParentID == pid {
			child := GetMenuTree(list,val.ID)
			node := dto.MenuTree {
				ID:        val.ID,
				ParentID:  val.ParentID,
				Type: string(val.Type),
				Name:      val.Name,
				Icon:      val.Icon,
				Path:      val.Path,
				Component: val.Component,
			}
			node.Children = child
			MenuTree = append(MenuTree,node)
		}
	}

	return  MenuTree
}


func MakeKeysInInSlice(haystack []string) func(needle string) bool {
	set := make(map[string]interface{})

	for _,e := range haystack{
		set[e] = struct {}{}
	}

	return func(needle string) bool {
		_,ok := set[needle]
		return ok
	}
}