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
6bafdddbfd6866b8482608e7708d139d309f40eb flag
6bafdddbfd6866b8482608e7708d139d309f40eb go/build
6bafdddbfd6866b8482608e7708d139d309f40eb io/ioutil
6bafdddbfd6866b8482608e7708d139d309f40eb os
6bafdddbfd6866b8482608e7708d139d309f40eb os/exec
6bafdddbfd6866b8482608e7708d139d309f40eb strings
```

Use this lock file to keep track of the dependencies. You don't need this lock file to compile your package, instead if your build fails for any reason you can use this file to compare the packages. [This](https://gist.github.com/oguzbilgic/a07ca257ac2af2f5c602) is a `.dondur.lock` generated for the [docker](https://github.com/dotcloud/docker) project.

#### External only mode

You can also use `-x` flag to just list the external packages.

```bash
$ dondur -x
```

This is the `.dondur.lock` file generated with the `-x` flag for [revel](http://github.com/robfig/revel).

```
4e0dc83f5a857e4d4f9455d1073eff284fdee117 code.google.com/p/go.net/websocket
08040c5a90632bd721465eb8ad74a8e61bd7bf95 github.com/howeyc/fsnotify
575cf31a8347a7889030f1f7fc4771be7dcd06fd github.com/robfig/config
f47995fbd5755034e17de0856f3eecfd2eff894f github.com/robfig/pathtree
6617b501e485b77e61b98cd533aefff9e258b5a7 github.com/streadway/simpleuuid
```

## Specs

`.dondur.lock` file is list of the packages imported by the project in the following format:

```
[git/hg commit hash] [package name]
```

*External* packages have `.` in the first part of their import paths'. For example:

```
github.com/oguzbilgic/socketio
code.google.com/p/go.net/websocket
```

## License

The MIT License (MIT)
