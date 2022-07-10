package statics

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"runtime"
)

func PathToHome() string {

	initialPath, error := user.Current()
	if error != nil {
		log.Fatal(error)
	}
	homeDirectory := initialPath.HomeDir

	return homeDirectory

}

func DeterminePath(homeDirectory string) string {

	var desktopPath string

	if runtime.GOOS == "linux" {
		desktopPath = string(homeDirectory + "/Desktop")
	}
	if runtime.GOOS == "windows" {
		desktopPath = string(homeDirectory + "\\Desktop")
	}

	return desktopPath

}

func CreateFolder(desktopFolder string) string {

	var pathToSites string

	files, err := ioutil.ReadDir(desktopFolder)
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir(desktopFolder)

	for _, file := range files {

		if file.Name() != "Sites" {
			os.Mkdir("Sites", os.ModePerm)
		}
	}

	if runtime.GOOS == "linux" {
		pathToSites = desktopFolder + "/Sites"
	}
	if runtime.GOOS == "windows" {
		pathToSites = desktopFolder + "\\Sites"
	}

	return pathToSites

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
