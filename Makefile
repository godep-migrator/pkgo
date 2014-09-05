fmt:
	go fmt ./...
test:
	godep go test -check.v
cov:
	rice embed-go; godep go test -cover -coverprofile=coverage.out
cover:
	go tool cover -html=coverage.out
dep:
	godep save -copy=false
