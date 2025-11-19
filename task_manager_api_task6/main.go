package main

import (
	"github.com/zaahidali/task_manager_api/data"
	"github.com/zaahidali/task_manager_api/router"
)

func main() {
	data.InitDB()
	r := router.SetupRouter()
	r.Run(":8080") // Start server
}
