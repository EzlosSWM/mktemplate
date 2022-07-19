package main

import (
	"fmt"
	app "mktemplate/app"
)

func main() {

	// Determines the path to the Desktop folder
	desktopFolder := app.DetermineDesktop()

	app.CreateSites(desktopFolder)
	sitesFolder := app.CreateSites(desktopFolder)

	// Moves into the created sites folder in the desktop & adds the index.html file
	app.CreateHTML(sitesFolder)

	// Creates folder for css & style.css
	app.CreateCSS(sitesFolder)
	app.CreateJS(sitesFolder)

	fmt.Printf("Complete, \"Sites\" folder created in %v.\n", sitesFolder)

}
