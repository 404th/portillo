package main

import (
	"fmt"

	"github.com/404th/portillo/config"
	"github.com/404th/portillo/config/data"
	"github.com/404th/portillo/pkg/mail"
)

func main() {
	// time.Sleep(2 * time.Second * time.Duration(1))
	// var waitGroup sync.WaitGroup
	for _, v := range data.Owners {
		// waitGroup.Add(1)
		mail.NewMailGen(config.MJ_APIKEY_PUBLIC, config.MJ_APIKEY_PRIVATE, v.Email, v.Name, data.Admin.Email, data.Admin.Name)
	}
	// waitGroup.Wait()

	fmt.Println("Here we go!")
}
