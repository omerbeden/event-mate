$docker create network --name event-mate-network


$docker run --rm -it --name dev-postgres -p 5432:5432 --network event-mate-network -e POSTGRES_PASSWORD=password -e POSTGRES_DB=test postgres

$docker run --rm -it --name dev-pgadmin  -p 80:80 --network event-mate-network -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' -e 'PGADMIN_CONFIG_LOGIN_BANNER="Authorised users only!"' -e 'PGADMIN_CONFIG_CONSOLE_LOG_LEVEL=10'  dpage/pgadmin4



##protoc  go sample
protoc --go_out=.  --go-grpc_out=.  ./proto/grpc/...

##dart sample
protoc -I ./proto/grpc/profile/v1 --dart_out=grpc:./frontend/event_mate_ui/pb ./proto/grpc/profile/v1/profile-service.proto