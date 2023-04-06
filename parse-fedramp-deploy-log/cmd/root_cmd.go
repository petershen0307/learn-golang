package cmd

import (
	"archive/zip"
	"log"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "app",
	}
	logFile string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&logFile, "file", "", "Could be like ./STG_Deployment.zip")
}

func Execute() error {
	err := rootCmd.Execute()
	parse()
	return err
}

func parse() {
	log.Println("input file=", logFile)
	// open zip file
	reader, err := zip.OpenReader(logFile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	// get all log file in a zip file
	for _, file := range reader.File {
		log.Println(file.Name)
	}
}
