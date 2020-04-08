default: main.go
	go run main.go

test: main_test.go main.go
	go test -v -cover ./...