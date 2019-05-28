package main

import (
	"os"

	"github.com/tommy-sho/golang-sandbox/web/payjp-vue/backend-api/infrastructure"
)

func main() {
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}
