1)installing a new go project

$ go mod init github.com/custordev/shortener-api

2)create a main.go file

touch main.go

3)running a go app

go run main.go

4)Hot reloading using the compile daemon
commands
go install github.com/githubnemo/CompileDaemon@latest
go get github.com/githubnemo/CompileDaemon@latest

Nb : the go.sum is used to see the packages in the app

compileDaemon -command="./shortener-api"
this adds the compile Daemon and also runs the build

5)Routing
gollia/mux for routing and also extracting params

command => go get -u github.com/gorilla/mux

6. Generating ids
   Google uuid

7. SQLITE
   modern package

Nb: difference between packages for installing and getting

go get => used in development
go install => used for installing

Gin Frame work

Gin is a web framework written in Go (Golang). It features a martini-like API with much better performance, up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

Gin works will work on routing ,params

go get -u github.com/gin-gonic/gin

gorom working as prisma
go get -u gorm.io/gorm

db postgres
go get -u gorm.io/driver/postgres

repo for go .env
https://github.com/joho/godotenv

go get github.com/joho/godotenv

jwt for go
jwt for authentication and authorization
go get github.com/golang-jwt/jwt/v5

bycrpt for go
for encrypting paswsword
go get golang.org/x/crypto/bcrypt

cors (cross origin resource share) for go
used
go get github.com/gin-contrib/cors

Nb creating a secret key
openssl rand -base64 32

steps 
we connect the environmental variables 
we connect the Db

Nb: c is going to get the data from the request in the controller using (c.shouldBindJSON)