.Phony: all, binary, lint, download
version := 1.0.1

binary: download lint
	@go build .

download:
	@go mod download
	ls ~/go/bin

lint:
	@~/go/bin/golint -set_exit_status ./...

docker:
	@docker build . --tag atif1996/podtester:$(version) --tag atif1996/podtester:latest
	@docker push atif1996/podtester:$(version)
	@docker push atif1996/podtester:latest
