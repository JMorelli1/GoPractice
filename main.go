package main

import (
	"net/http"

	"github.com/JMorelli1/GoPractice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
