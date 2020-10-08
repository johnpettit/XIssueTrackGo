# Compiler

FROM golang:1.15 as compiler
RUN mkdir /xissuetrackgo
WORKDIR /xissuetrackgo
COPY go.mod .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -ldflags '-s -w' -o /target .

# Target

FROM busybox:1.31.1

ENTRYPOINT ["/xissuetrackgo"]

EXPOSE 8088

COPY --from=compiler /target /xissuetrackgo
