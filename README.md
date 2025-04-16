# Golang-Github-Practices

## Git repo

git mod init github.com/hasanashik/Golang-Github-Practices

## Create go module named as git repo at the root of my project

go mod init github.com/hasanashik/Golang-Github-Practices

### To Run main.go

go run main.go
or
go run .

### output

```
hello world
```

# To Build the go program to binary

```
go build
```

### To run binary

./Golang-Github-Practices.exe
output

```
hello world
```

## Cross build/compile

See Current OS architechture
see environment variables
■ go env GOARCH GOOS
run one of these at the command line to build to a certain OS:
■ GOOS=darwin go build
■ GOOS=linux go build
■ GOOS=windows go build

## get puppy package

go get github.com/hasanashik/puppy@latest
go get github.com/hasanashik/puppy@1.0.1
