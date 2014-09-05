fmt:
	go fmt ./...
test:
	rice embed-go
	godep go test -check.v
	rm *.rice-box.go
cov:
	rice embed-go; godep go test -cover -coverprofile=coverage.out
cover:
	go tool cover -html=coverage.out
dep:
	godep save -copy=false
