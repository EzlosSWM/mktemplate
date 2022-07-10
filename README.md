# MKTEMPLATE

## MKTEMPLATE is a command line tool created using Go and is intended to create a folder dedicated to front-end web development. 

MKTEMPLATE creates a `Sites/` folder in the user's Desktop containing:

1. index.html 
2. css/
- css/style.css
3. app/ 
- app/app.js

The `css/style.css` and `app/app.js` is already linked to the index.html, making it a quick and seemless script to automate the beginning stages of front-end development.


# Installation
```bash 
git clone https://github.com/EzlosSWM/mktemplate.git
```

*Note: You can either modify the script or use it as is, or you can use the pre-compiled script within the builds folder*

# Usage 
```bash 
# Windows 
cd mktemplate\builds
./mktemplate.exe 

# Linux
cd mktemplate/builds/
./mktemplate
```

## End Result 
Once the folder is generated, you could open the `index.html` file on your favorite web browser and you'll see something simlar to: 

# Want to contribute?
Since this is the beta, I would eventually like to add: 
- Arguments to specify the desired path. 
- Arguments to specify the folder name. 

If you have any ideas or tips, you can contact me at: 
[Github](https://github.com/EzlosSWM) or
[Twitter](https://twitter.com/EzlosSWM).

# Known Issues
- Does not work for MacOS but will add it to the project very soon. 
