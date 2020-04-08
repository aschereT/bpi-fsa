default: main.go
	go build -ldflags="-s -w" -o bpi-fsa

test: main_test.go main.go
	go test -v -cover ./...

docker: main.go
	docker build . -t bpifsa

docker-test: main_test.go main.go
	docker build . -f Dockerfile.test