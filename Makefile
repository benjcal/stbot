build:
	go build -o stbot main.go

build-linux:
	GOOS=linux GOARC=x86_64 go build -o stbot main.go

clean:
	rm stbot