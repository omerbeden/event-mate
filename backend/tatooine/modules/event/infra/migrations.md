# How to run migrations from CLI

## Create

``` 
migrate create -ext sql -dir {path to migration dir} -seq {file name}
```

## Run
``` 
migrate -path {path} -s database {db path like postgres://postgres:password@localhost:5432/test?sslmode=disable} -verbose up
```
