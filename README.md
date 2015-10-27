umarugo
========


_himouto umaru language interpreter_

[![Build Status](https://travis-ci.org/135yshr/umrgo.png?branch=master)](https://travis-ci.org/135yshr/umrgo)

## Project Setup

_How do I, as a developer, start working on the project?_ 

```
go get github.com/r7kamura/gospel
go get github.com/135yshr/umrgo
```

## Testing

Using the [gospel](https://github.com/r7kamura/gospel) to the test framework.

### Unit Tests

1. `go test` or `go test ./...`

### Integration Tests

1. commit to the repository
2. Travis ci takes care of running the test automatically after a while.

## Deploying

### How to setup the deployment environment

- Please install the execution environment of the golang.
- Use the command go get, please get the gospel from github.

### How to deploy

```
go install
```


## Troubleshooting & Useful Tools

_error git push_

> remote: Invalid username or password.
> fatal: Authentication failed for 'https://github.com/135yshr/wspacego/'
> 
> - change push repository  
> `git remote set-url origin git@github.com:135yshr/umrgo.git`

## Change history

- relese version 1.0

## License
Copyright &copy; 2015 [135yshr](https://github.com/135yshr)
Licensed under the [Apache License, Version 2.0][Apache]

[Apache]: http://www.apache.org/licenses/LICENSE-2.0
[MIT]: http://www.opensource.org/licenses/mit-license.php
[GPL]: http://www.gnu.org/licenses/gpl.html
