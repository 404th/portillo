package main

import (
	"github.com/404th/portillo/config"
	"github.com/404th/portillo/pkg/email"
	"github.com/404th/portillo/pkg/mail"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("Error: ", err.Error())
		return
	}

	// email.NewMailGen(cfg.EmailPublicKey, cfg.EmailPrivateKey)

}
