package main

import (
	"fmt"

	"github.com/daymenu/goadmin/internal/config"
)

func main() {
	fmt.Println(config.AppDir())
}
