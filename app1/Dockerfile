FROM gcr.io/red-inspr/chimera:go-client AS build-env
WORKDIR /app
COPY go.mod go.mod
RUN go mod download
COPY . .
RUN go build -o main -tags musl main.go

FROM alpine
COPY --from=build-env /usr/lib/*kafka* /usr/lib/
COPY --from=build-env /app/main main
ENTRYPOINT ./main
