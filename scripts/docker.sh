$docker create network --name event-mate-network


$docker run -d --rm -p 5432:5432 --network event-mate-network postgres -h some-postgres -U postgres \
     -e POSTGRES_PASSWORD=mysecretpassword \

$docker run -d --rm -p 80:80 --network event-mate-network \
    -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' \
    -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' \
    -e 'PGADMIN_CONFIG_ENHANCED_COOKIE_PROTECTION=True' \
    -e 'PGADMIN_CONFIG_LOGIN_BANNER="Authorised users only!"' \
    -e 'PGADMIN_CONFIG_CONSOLE_LOG_LEVEL=10' \
    -d dpage/pgadmin4
