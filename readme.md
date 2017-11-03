# Lets vote
 
# About
This application lets you add polls and manage polls for your organisation.
# Download

# Compatibility

# Dependencies
The only dependency required to build the application is the go framework https://golang.org/

# Building the application
 1. Install, (if not already installed), the go framework https://golang.org/dl/
 2. run `go build` in the source directory
 3. (optional) rename the main binary to filecommander (or main.exe to filecommander if on windows)

# Configuration
example of `settings.json`
 ```
 {
   "binding": "0.0.0.0:8080"
 }
 
```

if binding is specified as  `0.0.0.0:8080` it will listen to all addresses  
if binding is specified as `192.168.0.100:80` it will listen to 192.168.0.100 at port 80

# Api

# Rebuild the front-end

## Preparations
 1. Download and install nodejs from https://nodejs.org (go for the LTS release if you are unsure which version to choose)
 2. Install webpack `npm install webpack -g`
 3. Install webpack development server `npm install webpack-dev-server -g`
 4. Install the additional dependencies run `npm install` in the frontend directory

## Developement
The front end is written using vue.js, vuex and webpack. There's no need to recompile the backend during development. 
The webpack development server will proxy the api calls to the backend once it is started.  
  

 1. Compile and start the backend application, let it listen to port 8080
 2. Start webpack-dev-server run `npm run dev` in the frontend directory
 3. browse to `localhost:8080` with your favourite browser
 4. _make your changes to the code_ and it will update in the browser as you save
 5. press `ctrl + c` to stop the dev server  

## Compile the front end code
 * Run `npm run build` to run webpack en embed the content into go
 * Compile the go source with the updated front end code
