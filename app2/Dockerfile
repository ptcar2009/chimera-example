FROM gcr.io/red-inspr/chimera:go-client
WORKDIR /app
COPY . .
RUN go mod download
CMD [ "go", "run", "app2/main.go" ]
