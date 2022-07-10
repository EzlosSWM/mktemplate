package main

import (
	"fmt"
	"log"
	"mktemplate/statics"
	"os"
)

func main() {

	// Determines the path to the Desktop folder
	homeDirectory := statics.PathToHome()
	desktopFolder := statics.DeterminePath(homeDirectory)

	statics.CreateFolder(desktopFolder)
	sitesFolder := statics.CreateFolder(desktopFolder)

	// Moves into the created sites folder in the desktop & adds the index.html file
	os.Chdir(sitesFolder)

	err := os.WriteFile("index.html", []byte(statics.Html), 0755)
	if err != nil {
		log.Fatalf("Unable to write file %v", err)
	}

	// Creates folder for css & style.css
	statics.CreateCSS(sitesFolder)
	statics.CreateJS(sitesFolder)

	fmt.Printf("Complete, \"Sites\" folder created in %v.\n", sitesFolder)

}
