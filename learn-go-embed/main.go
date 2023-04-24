package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type testScenarioTemplateParam struct {
	Name                  string
	RegionCode            string
	OSArch                string
	Environment           string
	IESDownloadURL        string
	LogCounterDownloadURL string
	AgentDownloadURLWin   string
	AgentDownloadURLMac   string
	AgentDownloadURLLinux string
	AGWFQDN               string
	Win32ProgramFilesDir  string
}

type renderLists struct {
	Tmpl  string
	Param testScenarioTemplateParam
}

//go:embed test/*
var scenarioShelf embed.FS

//-go:embed test/update_pii.json.gotmpl
//var test string

func main() {
	// go embed will embed the folder into binary
	scenarioShelf, err := fs.Sub(scenarioShelf, "test")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scenarios := []renderLists{
		{
			Tmpl: "test.json.gotmpl",
			Param: testScenarioTemplateParam{
				Name: "test",
			},
		},
	}
	for _, scenario := range scenarios {
		if contentFS, err := scenarioShelf.Open(scenario.Tmpl); err == nil {
			content, err := io.ReadAll(contentFS)
			if err == nil {
				fmt.Println(string(content))
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}
}
