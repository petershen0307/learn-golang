package main

import (
	"log"

	"github.com/petershen0307/learn-golang/parse-fedramp-deploy-log/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
