package statics

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"runtime"
)

var Path string

func DetermineDesktop() string {

	initialPath, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDirectory := initialPath.HomeDir

	// Creates the full path to the desktop
	if runtime.GOOS == "linux" {
		Path = string(homeDirectory + "/Desktop")
	}
	if runtime.GOOS == "windows" {
		Path = string(homeDirectory + "\\Desktop")
	}

	return Path

}

func CreateSites(Path string) string {

	var pathToSites string

	files, err := ioutil.ReadDir(Path)
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir(Path)

	for _, file := range files {

		if file.Name() != "Sites" {
			os.Mkdir("Sites", os.ModePerm)
		}
	}

	if runtime.GOOS == "linux" {
		pathToSites = Path + "/Sites"
	}
	if runtime.GOOS == "windows" {
		pathToSites = Path + "\\Sites"
	}

	return pathToSites

}

func CreateHTML(sitesFolder string) {

	os.Chdir(sitesFolder)

	err := os.WriteFile("index.html", []byte(Html), 0755)
	if err != nil {
		log.Fatalf("Unable to write file %v", err)
	}

}

func CreateCSS(sitesFolder string) {

	files, err := ioutil.ReadDir(sitesFolder)
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir(sitesFolder)

	for _, file := range files {

		if file.Name() != "css" {
			os.Mkdir("css", os.ModePerm)
		}
	}

	os.Chdir("css")

	fail := os.WriteFile("style.css", []byte(CssFile), 0755)
	if fail != nil {
		log.Fatalf("Unable to write file %v", fail)
	}

}

func CreateJS(sitesFolder string) {

	files, err := ioutil.ReadDir(sitesFolder)
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir(sitesFolder)

	for _, file := range files {

		if file.Name() != "app" {
			os.Mkdir("app", os.ModePerm)
		}
	}

	os.Chdir("app")

	fail := os.WriteFile("app.js", []byte(JS), 0755)
	if fail != nil {
		log.Fatalf("Unable to write file %v", fail)
	}

}
