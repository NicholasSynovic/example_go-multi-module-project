# `Go` Multi-Module Project Example

> A working `Go` project that builds multiple targets that use multiple internal
> and external dependencies

## Table of Contents

- [`Go` Multi-Module Project Example](#go-multi-module-project-example)
  - [Table of Contents](#table-of-contents)
  - [About](#about)
  - [How to Install](#how-to-install)
    - [Dependencies](#dependencies)
    - [Installation steps](#installation-steps)
  - [How to Run](#how-to-run)

## About

Programmers like to write *DRY* (don't repeat yourself) code. Programming
languages provide support to enable this, such as:

- `from` and `import` in Python,
- `#include` in C and C++, and
- `import` in Go.

These methods allow you to pull code from other *files*, *internal modules*, or
*external modules* of a project.

For small projects where only one executable is built, these methods are enough.
For larger projects that involve sub-modules or create multiple executables,
then enforcing *DRY* becomes difficult, particularly in `Go`.

To fix this, `Go` projects can leverage
[`workspaces` with `go.work`](https://go.dev/blog/get-familiar-with-workspaces).
There are several tutorials that explain how to do this, including:

- [The official `Go` tutorial](https://go.dev/doc/tutorial/workspaces),
- [Matt Boyle's post on Dev.to](https://dev.to/gophers/what-are-go-workspaces-and-how-do-i-use-them-1643),
- [Solomon Esenyi's post on LogRocket](https://blog.logrocket.com/go-workspaces-multi-module-local-development),
  and
- Sabbir Ahmed's work
  - [Medium Post](https://bysabbir.medium.com/go-workspaces-simplifying-multi-modular-projects-dc1a489302a)
  - [sabbir.dev](https://sabbir.dev/article/go-workspaces-simplifying-multi-modular-projects/)

However, a *complete* GitHub project that includes the `go.work` file, as well
as how to create multiple executables within a single project is not availible
to our knowledge. We thus present the first complete GitHub project to demo how
to leverage `go.work`, how to create multiple executables within this project,
and the supporting documentation to do so.

## How to Install

This code was tested on x86-64 Linux system.

### Dependencies

- `go`

### Installation steps

1. `git clone https://github.com/NicholasSynovic/example_go-multi-module-project`

## How to Run
