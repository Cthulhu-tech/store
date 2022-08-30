package main

import (
	"github.com/Cthulhu-tech/store/src/router"
	"github.com/Cthulhu-tech/store/src/utils/env"
)

func main() {

	env.Envload(".env")

	router.Handler()

}
