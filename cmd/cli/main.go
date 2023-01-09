package main

import (
	"fmt"
	"github.com/aasumitro/gowa/cmd/cli/console"
	"github.com/aasumitro/gowa/configs"
	"os"
)

func init() {
	configs.LoadEnv()

	configs.Instance.InitDbConn()
}

func main() {
	if err := console.Commands.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
