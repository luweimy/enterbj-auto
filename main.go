package main

import (
	"./client"
	"github.com/sirupsen/logrus"
)

func main() {
	c := client.New(client.Config{})
	info, err := c.GetPersonInfo("822F1CC8EB7142FBA6BDD7139C6AEA3E")
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(info)
}
