package main

import (
	"fmt"

	"github.com/gethinyan/enterprise/pkg/setting"
	"github.com/gethinyan/enterprise/routes"
)

func main() {
	listenAddr := fmt.Sprintf("0.0.0.0:%d", setting.Server.HTTPPort)
	router := routes.SetupRouter()
	router.Run(listenAddr)
}
