# go-fuck-yourself [![Build Status](https://travis-ci.org/adamryman/go-fuck-yourself.svg?branch=master)](https://travis-ci.org/adamryman/go-fuck-yourself)

```
go get github.com/adamryman/go-fuck-yourself
```

`go-fuck-yourself` is a `go` wrapper that will make your code build, forcefully.

### Do you have some code that will not build?

```
$ cat main.go
```
```
	package main

	import "fmt"

	// Sad code :(
	func main() {

		ss.Open()

		fmt.Println("What a time to be alive!")

		fmt.Printf()
		jfa;
		what
		even
		is
		this
	}
```
```
$ go build main.go
# command-line-arguments
./main.go:8: undefined: ss in ss.Open
./main.go:12: not enough arguments in call to fmt.Printf
./main.go:13: undefined: jfa
./main.go:14: undefined: what
./main.go:15: undefined: even
./main.go:16: undefined: is
./main.go:17: undefined: this
```

### Why not `go-fuck-yourself`?

```
$ go-fuck-yourself build main.go
FUCK: `ss.Open()`
FUCK: `fmt.Printf()`
FUCK: `jfa`
FUCK: `what`
FUCK: `even`
FUCK: `is`
FUCK: `this`
You are now a go developer
```
```
$ ./main
What a time to be alive!
```
```
$ cat main.go
	package main

	import "fmt"

	// Sad code :(
	func main() {


			fmt.Println("What a time to be alive!")

	}
```

And now, you too, can write go code.

Thanks [FuckItJs](https://github.com/mattdiamond/fuckitjs).

This software is in the public domain.
