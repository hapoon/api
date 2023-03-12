# syntax=docker/dockerfile:1
FROM golang:1.20.1-alpine3.17 AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 go build -o /hapoon-api .

FROM scratch AS release
COPY --from=build /hapoon-api /hapoon-api
CMD ["/hapoon-api"]
