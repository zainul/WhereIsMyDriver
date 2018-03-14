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
docker run --expose=5000 -p 5000:5000 -dit --restart unless-stopped --name whereismydriver --net="host" zainulmasadi/whereismydriver
```