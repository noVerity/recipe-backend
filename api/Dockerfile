FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /recipe-shard

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /recipe-shard /recipe-shard

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/recipe-shard" ]