Gin Microservice Example
========================

Version: 0.1.0

Example microservice using the Go [Gin](https://github.com/gin-gonic/gin) framework

### Main libraries
* [Gin](https://github.com/gin-gonic/gin) - API Router, controllers
* [Gom](https://github.com/mattn/gom) - Package management
* [Viper](https://github.com/spf13/viper) - Configuration
* [Stew](https://github.com/stretchr/stew) - Extends common Go objects providing better alternatives or wrappers

### Test libraries
* [Testify](https://github.com/stretchr/testify) - Testing framework
* [Mock](https://github.com/jarcoal/httpmock) - Mock underlying HTTP request/responses

## Create an API key

* Go to [OpenWeatherMap.org](http://openweathermap.org/api), sign up and copy the api key
* Uncomment the following line in config/development.toml

```
#app_id = "put-your-api-key-here"
```

* Replace the "put-your-api-key-here" with the api key from your OpenWeatherMap account


### Install go

```shell
$ brew install go
```

### Prepare ~/.bash_profile or terminal session for Go development

```shell
export GOPATH=$HOME/go
export GOROOT=/usr/local/opt/go/libexec
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

### Install this package

```shell
go get github.com/rbbll/gin-microservice
```

### Install gom

```shell
$ go get github.com/mattn/gom
```

### Download and install all dependencies/versions from Gomfile

```shell
$ cd $GOPATH/src/github.com/rbbll/gin-microservice
$ gom install
$ VIPER_CONFIG=../config gom test ./app
```

### Update dependencies (as needed)

```shell
$ rm -rf _vendor/
$ rm Gomfile.lock
$ gom lock
$ gom install
```

### Run tests

```shell
$ VIPER_CONFIG=../config gom test ./app
```

### Run tests (with logging)
```shell
$ VIPER_CONFIG=../config gom test -test.v ./app
```

### See test coverage
```shell
$ cd $GOPATH/src/github.com/rbbll/gin-microservice
$ mkdir -p coverage
$ VIPER_CONFIG=../config gom test -coverprofile=coverage/coverage.out ./app
$ VIPER_CONFIG=../config gom tool cover -html=coverage/coverage.out -o coverage/coverage.html
$ open coverage/coverage.html
```

### Start the application (development mode)

```shell
$ gom run main.go
```

### Run commands from the project folder (Standalone binary)

```shell
$ gom build
$ VIPER_ENV=development VIPER_CONFIG=./config ./gin-microservice
```
