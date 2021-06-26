# Overview

This project shows "import cycle" problem in golang and its solution.

- **Problem:** [play0](play0/) project shows "import cycle" problem

- **Solution:** [play1](play1/) project shows how to solve "import cycle" problem in play0 project

*Seach **IMP** in code to know which code creates "import cycle" problem.*

# Run

Run **play0**, it gives "import cycle not allowed" error.

```
$  go-import-cycle git:(master) ✗ cd play0
$  play0 git:(master) ✗ tree
.
├── foo
│   ├── bar
│   │   └── bar.go
│   └── foo.go
├── go.mod
└── main.go

2 directories, 4 files
$  play0 git:(master) ✗ go run main.go 
package command-line-arguments
        imports play0/foo
        imports play0/foo/bar
        imports play0/foo: import cycle not allowed
```

Run **play1**, it runs successfully and print output.

```
$  go-import-cycle git:(master) ✗ cd play1 
$  play1 git:(master) ✗ tree
.
├── foo
│   ├── bar
│   │   └── bar.go
│   └── foo.go
├── foobar
│   └── foo
│       ├── bar
│       │   └── bar.go
│       └── foo.go
├── go.mod
└── main.go

5 directories, 6 files
$  play1 git:(master) ✗ go run main.go 
play1
hello_from_foo
hello_from_bar
```

Note, you can structure **play1** as below with only **foobar** package to hold all interfaces.
```
$  play1 git:(master) ✗ tree 
.
├── foo
│   ├── bar
│   │   └── bar.go
│   └── foo.go
├── foobar
│   ├── bar.go
│   └── foo.go
├── go.mod
└── main.go
```

# References

- OK Let’s Go: Three Approaches to Structuring Go Code - [link](https://www.perimeterx.com/tech-blog/2019/ok-lets-go/)

