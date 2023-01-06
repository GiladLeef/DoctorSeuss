FROM golang:latest

LABEL maintainer="Hugo Lageneste <hugo@github.com/GiladLeef/DoctorSeuss>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go test
RUN go build -o main .
EXPOSE $PORT

CMD ["./main", "-port=$PORT"]