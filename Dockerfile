FROM golang:1.23

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go install github.com/codegangsta/gin@latest
CMD ["gin", "-i", "-a", "8080", "run", "main.go"]