# syntax=docker/dockerfile:experimental
FROM golang:1.14.3-alpine3.11 AS common
WORKDIR /go/src/app
COPY . .

FROM common AS build-api
RUN CGO_ENABLED=0 go build -o /hapoon-api cmd/hapoon/main.go

FROM scratch AS release
COPY --from=build-api /hapoon-api /
CMD ["/hapoon-api"]
