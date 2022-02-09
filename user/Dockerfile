FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /user-service

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /user-service /user-service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/user-service" ]