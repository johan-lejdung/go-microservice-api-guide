# REST API

Read the accompanying article here: https://medium.com/@johanlejdung/a-mini-guide-build-a-rest-api-as-a-go-microservice-together-with-mysql-fc203a6411c0

Included here are a `docker-compose` file that can be used to spin up a MySQL database in docker for the porpose of running this code.

Run the code:
```
go run main.go
```

Try inserting a value with
```
curl -XPOST localhost:8080/endpoint -v
```

Then fetching a value with
```
curl localhost:8080/endpoint/1 -v
```
