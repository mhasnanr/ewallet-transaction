package main

import (
	"fmt"

	"github.com/mhasnanr/e-wallet/bootstrap"
)

func main() {
	bootstrap.SetupConfig()
	fmt.Println(bootstrap.GetEnv("DB_NAME", ""))
}