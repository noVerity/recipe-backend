FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /gateway

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /gateway /gateway

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/gateway" ]