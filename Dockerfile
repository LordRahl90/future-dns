 FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o dns ./cmd/


FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /app/dns dns

EXPOSE 8080

ENTRYPOINT [ "./dns" ]
