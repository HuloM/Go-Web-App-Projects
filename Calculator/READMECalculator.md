# How to use this app
> Download files into one folder
> run following command in either a bash or zsh shell
```
$ GOARCH=wasm GOOS=js go build -o lib.wasm main.go
$ go run server.go
```
then connect to localhost:8080
web app should be up and running
