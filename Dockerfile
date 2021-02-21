FROM golang:1.16 as builder

# modules
WORKDIR $GOPATH/src/github.com/shanehowearth/argyle
ADD . $GOPATH/src/github.com/shanehowearth/argyle

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# build time
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/api cmd/main.go cmd/routes.go

FROM scratch
WORKDIR /root/
COPY --from=0 /go/bin/api .

# run the rest server
ENTRYPOINT ["./api"]
