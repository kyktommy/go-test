package main

import (
	"fmt"

	"github.com/kyktommy/go-test/pkg/log"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("main")

	logger.Log2("go-test logger")

	viper.SetDefault("ContentDir", "content")
}
