# syntax=docker/dockerfile:experimental
FROM golang:1.20.1-alpine3.17 AS build
WORKDIR /go/src/app
COPY . .

FROM build
RUN CGO_ENABLED=0 go build -o /hapoon-api cmd/hapoon/main.go

FROM scratch AS release
COPY --from=build /hapoon-api /
CMD ["/hapoon-api"]
