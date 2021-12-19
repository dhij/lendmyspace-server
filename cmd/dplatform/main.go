package main

import (
	"github.com/davidhwang-ij/study-platform/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run(":8080")
}
