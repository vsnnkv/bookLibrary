FROM golang:1.18-alpine

COPY go.mod ./
COPY go.sum ./
ENV GOPATH=/

RUN go mod download

COPY ./ ./

RUN go build -o book-app ./main.go

EXPOSE 8080

CMD [ "./book-app"]