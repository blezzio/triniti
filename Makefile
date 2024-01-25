test:
	go test -coverprofile=./test_profile ./... && go tool cover -html=./test_profile && unlink ./test_profile