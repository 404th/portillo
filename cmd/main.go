package main

import (
	"github.com/404th/portillo/config"
	"github.com/404th/portillo/config/data"
	"github.com/404th/portillo/pkg/email"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("Error: %s", err.Error())
		return
	}

	for _, v := range data.Owners {
		email.NewMailGen(cfg.EmailPublicKey, cfg.EmailPrivateKey, v.Email, v.Name, data.Admin.Email, data.Admin.Name)
	}
}
