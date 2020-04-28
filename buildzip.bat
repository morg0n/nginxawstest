set GOOS=linux
go build -o main main.go
build-lambda-zip -o main.zip main 