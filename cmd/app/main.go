package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := app(); err != nil {
		logrus.Panic(err)
		os.Exit(0)
	}
}
