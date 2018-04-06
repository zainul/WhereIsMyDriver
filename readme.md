# Tech Stack

## Languange (Go)
GO is prgramming langsung that type safety and we can play arround the memory management but the code is simple 

## Framework (Iris)
iris is simple and high throughput framework for golang

## Database (Mysql)
mysql is database can support ```ACID``` and leightweight

# Installation

## Installation on K8s

### Pre requirement (on Kubernetes)
- install k8s
- install network

### Installation Step:
- provide the build image in docker 
- provide the yaml-service of the service
- provide the yaml-deployment of the service

set k8s service of the service
```
kubectl create -f service.yaml
```

set k8s deployment of the service
```
kubectl create -f deployment.yaml 
```

- make proxy pas with nginx or HAproxy


## Installation on Docker

### Pre requirement (on Kubernetes)
- install docker

### Installation Step:
- provide the build image in docker 

run the docker 
```
docker run --expose=3000 -p 3001:3000 -dit --restart unless-stopped --name whereismydriver --net="host" zainulmasadi/whereismydriver
```
### Installation app in local

install iris

```
go get -u github.com/kataras/iris
```

run this command to install dependency
```
go get ./...
```
### Test

for integration (http) test

```
cd $GOPATH/src/WhereIsMyDriver/router && go test -v
```

for unit test

```
cd $GOPATH/src/WhereIsMyDriver/model && go test -v
```

for load test

requirement for load test:

1. install load test tool (mgun)

```
https://github.com/byorty/mgun
```

and then follow the instruction

```
cd /path/to/gopath
export GOPATH=/path/to/gopath/
export GOBIN=/path/to/gopath/bin/
go get github.com/byorty/mgun
go install src/github.com/byorty/mgun/mgun.go
```

after this run the script

go to your path

```
./bin/mgun -f src/WhereIsMyDriver/mgun_test.yaml
```
