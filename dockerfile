FROM golang:1.20

ENV DOCKER_API_VERSION=1.41
COPY . /app
WORKDIR /app
RUN go mod tidy
CMD ["go", "run", "main.go"]
