default: main.go
	go run main.go

test: main_test.go main.go
	go test -v -cover ./...

docker: main.go
	docker build . -t bpifsa && docker run -t bpifsa

docker-test: main_test.go main.go
	docker build . -f Dockerfile.test