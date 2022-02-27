$docker create network --name event-mate-network


$docker run --rm -it --name dev-postgres -p 5432:5432 --network event-mate-network -e POSTGRES_PASSWORD=password -e POSTGRES_DB=test postgres

$docker run --rm -it --name dev-pgadmin  -p 80:80 --network event-mate-network -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' -e 'PGADMIN_CONFIG_LOGIN_BANNER="Authorised users only!"' -e 'PGADMIN_CONFIG_CONSOLE_LOG_LEVEL=10'  dpage/pgadmin4


docker run --rm -it --name dev-postgres -e POSTGRES_DB=test -e POSTGRES_PASSWORD=Pass2020! -p 5432:5432 postgres