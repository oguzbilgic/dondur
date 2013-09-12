# Dondur

freeze your Go dependencies with ease

## Installation

If you have `go tools` installed on your system, use the command bellow to get the `dondur`:

```bash
$ go get github.com/oguzbilgic/dondur
```

## Usage

Dondur basically generates `.dondur.lock` file for your go package. This lock file consists of all the imported packages and their current version control hashes. In your package's directory, run the command bellow.

```bash
$ dondur
```

This will generate `.dondur.lock` file in your package's directory. Go ahead and commit this lock file to your source control system. Here is the lock file generated for this package.

```
go/build - 6bafdddbfd6866b8482608e7708d139d309f40eb
io/ioutil - 6bafdddbfd6866b8482608e7708d139d309f40eb
os - 6bafdddbfd6866b8482608e7708d139d309f40eb
os/exec - 6bafdddbfd6866b8482608e7708d139d309f40eb
strings - 6bafdddbfd6866b8482608e7708d139d309f40eb
```

Use this lock file to keep track of the dependencies. You don't need this lock file to compile your package, instead if your build fails for any reason you can use this file to compare the packages.

#### External only mode

You can also use `-x` flag to just list the external packages.

```bash
$ dondur -x
```

This is the `.dondur.lock` file generated with the `-x` flag for [socketio go package](http://google.com).

```
code.google.com/p/go.net/websocket - 4e0dc83f5a857e4d4f9455d1073eff284fdee117
```

## Specs

`.dondur.lock` file is list of the packages imported by the project in the following format:

```
[ package name ] - [ git/hg commit hash ]
```

*External* packages have `.` in the first part of their import paths'. For example:

```
github.com/oguzbilgic/socketio
code.google.com/p/go.net/websocket
```

## License

The MIT License (MIT)
