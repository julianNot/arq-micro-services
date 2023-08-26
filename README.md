# golang-gorm-api

# Golang Version
GO 1.20
 
https://go.dev/dl/

#  Docker Instructions
## Download Postgres image

```
docker pull postgres
```

## Run postgres container with environment variables (CREATE)

```
docker run --name [nameContainer] -e POSTGRES_USER=[user] -e POSTGRES_PASSWORD=[password] -e POSTGRES_DB=[nameDB] -p 5432:5432 -d postgres
```

## Run the Postgres container after creating it
```
docker start [nameContainer]
```

## Access the container

```
docker exec -it [nameContainer] bash
```

## Access the Postgres in Container

```
psql -U [user] --password --db [nameDB]
```

# Dependencies
* Mux
* GORM - ORM
* Driver Postgres - GORM
* AIR (https://github.com/cosmtrek/air)
* Godotenv

# Run Develop environment
Run the command in the root of project
```
air
```

# Run Production environment
Compile
```
go build main.g
```
```
./main
```