# gin-jwt-react-query-auth

 - gin-jwt https://github.com/appleboy/gin-jwt
 - react-query-auth https://github.com/alan2207/react-query-auth

## Backend(Go)

### initial

```
docker-compose build
docker-compose run --rm web air init
docker-compose run --rm web go mod tidy
```

### up

```
docker-compose up -d
```

### add package

```
cd src
go mod tidy
```

when installed, go.mod and go.sum are updated.

### ref 

https://zenn.dev/hrs/articles/go-gin-air-docker

## Frontend(React)

```
cd front
npm install
npm start
```

http://localhost:3000/
