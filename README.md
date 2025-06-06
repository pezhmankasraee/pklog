# pklog
__pklog__ providing log files on syslog and live on stdout written in Golang. This package is written in `Golang`. This logging package is able to generate two types of log information both on:
* syslog
* stdout

## Badges
![golang version](https://img.shields.io/badge/golang-v1.24.3-lightblue)

## Author
Pezhman Kasraee [pezhmankasraee@gmail.com](pezhmankasraee@gmail.com).

## Download the package
```bash
go get -u github.com/pezhmankasraee/pklog
```

## Example
Before starting the application, you shoud set the environmental variable PKLOG. If the application is going to be run on development mode, you should set it as follows:
```bash
export PKLOG="DEVELOPMENT"
```
or if the application is going to be run on server, the environment variable should be set as follows:
```bash
export PKLOG="PRODUCTION"
```
Then you should write your code as follows:
```go
package main

import "github.com/pezhmankasraee/pklog"

func main() {
	pklog.CreateLog(pklog.Error, "Hello, Error")
}
```
if you run the application with `PKLOG="DEVELOPMENT"`, you will see the result on `stdout`, or if you run it with `PKLOG="PRODUCTION"` the result will be reflected on /var/log/syslog. You can see the result by:
```bash
$ cat /var/log/syslog
```

## Versions
* Version 1.x.x
	module: go-pklog

* Version 2:

	release: 6 June 2025

	NOTE: this release is backward compatible with version 1.x.x

	major change:

		1. module name changed to `pklog`.
		2. License changed to BSD3.
