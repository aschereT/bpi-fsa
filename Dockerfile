FROM golang:1.13-alpine AS build

WORKDIR /bpifsa

ADD . .
RUN go build -ldflags="-s -w" -o bpi-fsa

FROM alpine:3.11.5 AS final

COPY --from=build /bpifsa/bpi-fsa /bin/bpi-fsa

ENTRYPOINT [ "/bin/bpi-fsa" ]