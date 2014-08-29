fmt:
	go fmt ./...
test:
	godep go test -check.v
dep:
	godep save -copy=false
