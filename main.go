package main

import (
	"fmt"
	statics "mktemplate/app"
)

func main() {

	// Determines the path to the Desktop folder
	desktopFolder := statics.DetermineDesktop()

	statics.CreateSites(desktopFolder)
	sitesFolder := statics.CreateSites(desktopFolder)

	// Moves into the created sites folder in the desktop & adds the index.html file
	statics.CreateHTML(sitesFolder)

	// Creates folder for css & style.css
	statics.CreateCSS(sitesFolder)
	statics.CreateJS(sitesFolder)

	fmt.Printf("Complete, \"Sites\" folder created in %v.\n", sitesFolder)

}
