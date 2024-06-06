# `Go` Multi-Module Project Example Project Tutorial

> The tutorial on how to leverage `go work` to create multi-module projects

## Table of Contents

- [`Go` Multi-Module Project Example Project Tutorial](#go-multi-module-project-example-project-tutorial)
  - [Table of Contents](#table-of-contents)
  - [About](#about)
  - [Tutorial](#tutorial)
    - [Notes About Our SE Practice](#notes-about-our-se-practice)
    - [Creating A Multi-Module Project](#creating-a-multi-module-project)
    - [Creating Sub-Modules](#creating-sub-modules)
    - [Adding "Visible" and "invisible" Code](#adding-visible-and-invisible-code)
    - [Creating The Executable](#creating-the-executable)
    - [Running the Executable](#running-the-executable)

## About

This `README.md` file is meant to provide the software engineering knowledge
necessary to facilitate `go work` usage. It is meant to add to existing work on
this, while providing concrete examples.

## Tutorial

This tutorial is structured to provide:

- Clean, step by step instructions for creating a multi-module project, and
- Code that can be copied into a terminal to follow along.

It does not:

- Provide `Go` code in the code block to insert into your own files.

Each code section builds on one another. For example:

```shell
# Code Section 1
mkdir temp
cd temp
```

and

```shell
# Code Section 2
cd ..
rm -rf temp
```

`Code Section 2` builds on `Code Section 1`. Without `Code Section 1`,
`Code Section 2` won't work.

### Notes About Our SE Practice

This project and the tutorial were written by a small team. If you are working
in a large team, then adapt the following notes to fit your organization.

We follow a consistent style throughout this tutorial. In practice, folder,
modules, file, and function names can be renamed to meet style standards (with
the exception of capital vs lower case function names; see
[this section](#adding-visible-and-invisible-code)).

### Creating A Multi-Module Project

1. Create an empty `src` directory
1. `cd` into the empty directory
1. `go work init` to create an empty `go.work` file

```shell
mkdir src
cd src
go work init
```

### Creating Sub-Modules

1. Inside the `src` directory, create a directory to store your module code
   (`string_module`)
1. Inside the `src` directory, create a directory to store your executable code
   (`cmd`)
1. For each executable, create a directory within the `src/cmd` directory
1. For each newly created directory (except the top level of the `cmd`
   directory), execute `go mod init PROJECT/MODULE` where `PROJECT` is the name
   of your project and `MODULE` is the name of your module
1. For each newly created directory (except the top level of the `cmd`
   directory), execute `go work use .`
   
```shell
mkdir string_module
mkdir cmd
mkdir cmd/hello_world

cd string_module && go mod init example-project/string_module && go work use . && cd ..
cd cmd/hello_world && go mod init example-project/hello_world && go work use . && cd ../..
```

Here `PROJECT` is `example-project` and each `MODULE` is the name of the
directory.

### Adding "Visible" and "invisible" Code

1. In `string_module`, create a file called `func.go`
1. In `func.go`, create two functions:
   1. One that starts with a capital letter (e.g., `func HelloWorld()`), and
   1. One that starts with a lower case letter (e.g, `func addPunctuation()`)

```shell
cd string_module
touch func.go
# Implement Go code
```

Here our "Visible" code starts with a capital letter; this code can be accessed
by files that import the module. Our "invisible" code starts with a lower case
letter; this code can only be accessed by files within the module.

### Creating The Executable

1. Create `main.go` in `cmd/hello_world`
1. Import `string_module` as `sm`
1. Implement `main` function that uses `sm.HelloWorld()`

```shell
touch cmd/hello_world/main.go
# Implement main function
```

### Running the Executable

1. Run `go run` on the `main.go` code

`go run cmd/hello_world/main.go`
