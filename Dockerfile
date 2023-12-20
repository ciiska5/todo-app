FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

#install psql for wait-for-postgres.sh
RUN apt-get update
RUN apt-get -y install postgresql-client

#make the wait-for-postgres.sh file executable
RUN chmod +x wait-for-postgres.sh

#build todo-app
RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]