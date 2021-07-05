# Calculator app
this app is my first attempt at a go program, created from a tutorial found: https://tutorialedge.net/golang/go-webassembly-tutorial/ 
you can find most of the code I used their minus some changes I made to both the [index page](./index.html) and [main go file](./main.go) 

# How to use this app
> Download files into one folder
> run following command in either a bash or zsh shell
```
$ cd (path of directory of project files)/
$ GOARCH=wasm GOOS=js go build -o lib.wasm main.go
$ go run server.go
```
then connect to localhost:8080
web app should be up and running
