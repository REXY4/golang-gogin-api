#gogin

```
go mod init
```


```
go run main.go
```

created db 
```
createdb -h localhost -p 5432 -U rexy4 gogin --password
```
run with nodemon
```
nodemon --exec go run main.go --signal SIGTERM
```