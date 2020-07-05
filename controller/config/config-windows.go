package config

// document: https://github.com/winsw/winsw

import (
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
)

type WindowsConfig struct {
}

type WinswService struct {
	XMLName          xml.Name `xml:"service"`
	Id               string   `xml:"id"`
	Name             string   `xml:"name"`
	Description      string   `xml:"description"`
	Executable       string   `xml:"executable"`
	WorkingDirectory string   `xml:"workingdirectory"`
	Argument         []string `xml:"argument"`
	Logmode          string   `xml:"logmode"`
}

func (config *WindowsConfig) Config(homeDir, executable string, args ...string) error {
	service := WinswService{
		Id:               "MyDemoServer",
		Name:             "MyDemoServer",
		Description:      "A demo service run on Windows operation system.",
		Executable:       executable,
		Logmode:          "rotate",
		WorkingDirectory: homeDir,
	}
	for _, arg := range args {
		service.Argument = append(service.Argument, arg)
	}
	bs, err := xml.MarshalIndent(service, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(homeDir, "service.xml"), bs, 0666)
}
