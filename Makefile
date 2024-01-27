test:
	go test -coverprofile=./test_profile ./... && go tool cover -html=./test_profile && unlink ./test_profile

build:
	GOOS=windows GOARCH=386 go build -o bin/triniti-386.exe main.go && \
    GOOS=windows GOARCH=amd64 go build -o bin/triniti-amd64.exe main.go && \
    GOOS=darwin GOARCH=amd64 go build -o bin/triniti-amd64-darwin main.go && \
    GOOS=linux GOARCH=386 go build -o bin/triniti-386-linux main.go && \
    GOOS=linux GOARCH=amd64 go build -o bin/triniti-amd64-linux main.go