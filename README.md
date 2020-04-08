# bpi-fsa

## Running

### With Golang installed

See https://golang.org/doc/install for installation instructions.

Once Golang is installed, run `make` to build the binary, then invoke with `./bpi-fsa`.

Run tests/coverage with `make test`

### Docker

See https://www.docker.com/get-started for installation instructions.

Run `make docker` to build the image, then use `docker run -t bpifsa` to run the container

Run `make docker-test` to run the tests inside a Docker container.