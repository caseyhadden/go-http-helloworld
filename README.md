A simple Go program that listens for HTTP GET requests on a port and responds with
hello world.

It is ultimately used as part of a basic, demonstration BOSH release.

In order to build and run this project, you'll need a valid environment for building Go
code as a prerequisite.

Cloning and building the project:

``` bash
$ git clone git://github.com/caseyhadden/go-http-helloworld.git
$ cd go-http-helloworld
$ go build
```

Options for running:

``` bash
$ ./go-http-helloworld         # will listen on port 8080
$ ./go-http-helloworld 9999    # will listen on port 9999
```
