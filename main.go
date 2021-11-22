package main

import (
	"golang_api/config"
	"golang_api/router"
)

func main()  {

	r,db := router.RouteMap()
	r.Static("/static","./static")
	defer config.CloseDatabaseConnection(db)
	err := r.Run()
	if err != nil {

		return
	}
}
