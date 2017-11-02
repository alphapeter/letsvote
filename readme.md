# file commander
 
*Simple Web File Commander*

# About
A simple, portable web file commander implemented in the Go language. 
It compiles into one single binary which bundles all HTML, javascript etc. 
The protocol for the API is JSON RPC 2.0 through HTTP posts.
You can define virtual file system roots to make certain parts of the file system accessible.  

# Download
https://github.com/alphapeter/filecommander/releases

# Compatibility
The front end will only run on modern browsers, such as chrome, firefox, opera or edge, **internet explorer will not work**
# Dependencies
The only dependency required to build the application is the go framework https://golang.org/

# Building the application
 1. Install, (if not already installed), the go framework https://golang.org/dl/
 2. run `go build` in the source directory
 3. (optional) rename the main binary to filecommander (or main.exe to filecommander if on windows)

# Configuration
Before you can utilize the file commander, you have to define your (virtual) file system roots.
example of `settings.json`
 ```
 {
   "roots": [
     {
       "name": "temp",
       "path": "/tmp"
     },
     {
       "name": "incoming",
       "path": "/var/incoming"
     }
   ],
   "binding": "0.0.0.0:8080"
 }
 
```

example of `settings.json` for windows
 ```
 {
   "roots": [
     {
       "name": "temp",
       "path": "c:/temp"
     },
     {
       "name": "incoming",
       "path": "c:/incoming"
     }
   ],
   "binding": "0.0.0.0:8080"
 }
 
```

if binding is specified as  `0.0.0.0:8080` it will listen to all addresses  
if binding is specified as `192.168.0.100:80` it will listen to 192.168.0.100 at port 80

**file location**  
* the default search path is './settings.json'  
* to load settings file from another location, use the argument `--settings <path-to-settings>/<filename>.json` when starting the application  

# Protocol

## JSON RPC 2.0
For complete specification of the JSON RPC protocol, please visit http://www.jsonrpc.org/specification

## Copy (cp)
Copy a file to an other file

Example of command:
```
{
  "jsonrpc":"2.0",
  "method": "cp",
  "params": ["private/animals/cat.jpg", "public/animals/cat.jpg"],
  "id": "3"
}
```
Example of successful response:
```
{
    "jsonrpc":"2.0",
    "result":null,
    "id":"3"
}
```
Example of unsuccessful response:
```
{
    "jsonrpc":"2.0",
    "id":"3",
    "error":
    {
        "code":-32603,
        "message":"open /temp/animals/cat.jpg: The system cannot find the file specified."
    }
}
```

## List roots (df)
Lists available (virtual) file system roots by name  

Example of command:
```
{ 
  "jsonrpc":"2.0",
  "method": "df",
  "params": [],
  "id": "3"
}
```
Example of successful response:
```
{
    "jsonrpc":"2.0",
    "result":["incoming","temp"],
    "id":"3"
}
```
## List files (ls)
Lists all files and directories for a certain path

Example of command:
```
{
  "jsonrpc":"2.0",
  "method": "ls",
  "params": ["incoming/public"],
  "id": "3"
}
```
Example of successful response:
Note: Type d: Directory, Type f: File
```
{
    "jsonrpc":"2.0",
    "result":[{"Type":"d","Name":"dir1"},{"Type":"f","Name":"file1.txt"},{"Type":"f","Name":"file2.txt"}]
}
```
## Make directory (mkdir)
Creates a directory

Example of command:
```
{
  "jsonrpc":"2.0",
  "method": "mkdir",
  "params": ["incoming/animals", "cats"],
  "id": "3"
}
```
Example of successful response:
```
{
    "jsonrpc":"2.0",
    "result":null,
    "id":"3"
}
```
## Move (mv)
Moves/renames a file or directory
Note. A file cannot be moved between two phycically different drives

Example of command:
```
{
  "jsonrpc":"2.0",
  "method": "mv",
  "params": ["incoming/public/dog.jpg", "public/animals/dog.jpg"],
  "id": "3"
}
```
Example of successful response:
```
{
    "jsonrpc":"2.0",
    "result":null,
    "id":"3"
}
```
## Delete (rm)
Deletes a file  

Example of command:
```
{ 
  "jsonrpc":"2.0",
  "method": "rm",
  "params": ["incoming/public/dog.jpg"],
  "id": "3"
}
```
Example of successful response:
```
{
    "jsonrpc":"2.0",
    "result":null,
    "id":"3"
}
```

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
 
# Work in progress (whats next)
 **TODO**
 * implement key navigation
 * implement demo 
 * make precompile binaries available
 * code cleanup
