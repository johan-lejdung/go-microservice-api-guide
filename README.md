# REST API

## Before cloning

Read the accompanying article here: https://medium.com/@johanlejdung/a-mini-guide-build-a-rest-api-as-a-go-microservice-together-with-mysql-fc203a6411c0

Change the module name if you are cloning it to a new location.

```
go mod init <your_module_name>
```

If you are using a private repository make sure to set GOPRIVATE with:

```
go env -w GOPRIVATE=github.com/repoURL/private-repo
```

## Setup

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

## Further reading

If you are interested in optimal folder structure, have a look at this [github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout).
