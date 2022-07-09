# REST API for test case backend in Majoo

## Documentation
[Postman Documentation](https://documenter.getpostman.com/view/12132212/UzJPLEp5)

 [Swagger/OPEN API Documentation](https://editor.swagger.io/)
 1. Please open link https://editor.swagger.io/
 2. Copy all code in file `apispesification.json` and paste into web swagger on top

## Run 
```
go run main.go
```

## Use docker with docker-compose
Docker is used for run container database MySQL.
```
docker-compose -f docker-compose.dev.yml up -d
```