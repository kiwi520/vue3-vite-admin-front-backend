package main

import (
	"golang_api/config"
	"golang_api/router"
)

func main()  {

	r,db := router.RouteMap()
	defer config.CloseDatabaseConnection(db)
	err := r.Run()
	if err != nil {

		return
	}
}
