# ober
Simple Golang based file server

## Compiling
You should have the [Golang](http://golang.org) compiler.
After installing that, you simply navigate to the directory in which the ober.go file is located. And you invoke the following command.
```
$ go build ober.go
```

## Usage
If all was successful, you can now invoke the file server like so:
```
$ ./ober --file=<filename>
```
After which it will serve the file at http://localhost:3232/

## Optional parameters
You can list other parameters ( for example --port ) by invoking:
```
$ ./ober --help
```
