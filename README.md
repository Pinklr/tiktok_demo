# tiktok_demo


## Quick Start
### 0.Modify ServerAddr
Change the constant `ServerAddr` in file`pkg/constants` ServerAddr to local ip address.
### 1.Setup Basic Dependence
```shell
docker-compose up
```

### 2.Run User RPC Server
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 3.Run Video RPC Server
```shell
cd cmd/video
sh build.sh
sh output/bootstrap.sh
```
### 4.Run Interact RPC Server
```shell
cd cmd/interact
sh build.sh
sh output/bootstrap.sh
```
### 5.Run API Server
```shell
cd cmd/api
go run .
```