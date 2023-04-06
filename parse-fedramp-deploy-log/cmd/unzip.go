package cmd

import (
	"archive/zip"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

var unzipCMD = &cobra.Command{
	Use:   "unzip [zip file path]",
	Short: "unzip file to current folder",
	Args:  cobra.ExactArgs(1),
	Run:   unzip,
}

// flags
var (
	targetDir string
)

func init() {
	workingDir, _ := os.Getwd()
	unzipCMD.PersistentFlags().StringVarP(&targetDir, "target", "t", workingDir, "Specified target directory")
}

func unzip(cmd *cobra.Command, args []string) {
	zipFilePath := args[0]
	log.Println("input zip file=", zipFilePath)
	// open zip file
	zipReader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// check target dir existed
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		err = os.Mkdir(targetDir, 0777)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// check target is a directory
	if f, _ := os.Stat(targetDir); !f.IsDir() {
		log.Fatalln(targetDir, " is not a directory")
	}

	const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	var re = regexp.MustCompile(ansi)
	// get all log file in a zip file
	for _, fileInZIP := range zipReader.File {
		fileReader, err := fileInZIP.Open()
		if err != nil {
			log.Printf("file name=%v, err=%v\n", fileInZIP.Name, err)
			continue
		}
		bufScanner := bufio.NewScanner(fileReader)
		outputFile, err := os.OpenFile(filepath.Join(targetDir, fileInZIP.Name), os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Printf("open file=%v, err=%v\n", filepath.Join(targetDir, fileInZIP.Name), err)
			continue
		}
		for bufScanner.Scan() {
			text := bufScanner.Text()
			if _, err = outputFile.WriteString(re.ReplaceAllString(text, "") + "\n"); err != nil {
				log.Printf("write file=%v, err=%v\n", filepath.Join(targetDir, fileInZIP.Name), err)
				continue
			}
		}
		if err = fileReader.Close(); err != nil {
			log.Println("Close file=", fileInZIP.Name, " Error=", err)
		}

		if err = outputFile.Close(); err != nil {
			log.Println("Close file=", fileInZIP.Name, " Error=", err)
		}
	}
}
