FROM golang:latest
WORKDIR /app
COPY . .
RUN export GOPROXY=https://goproxy.io,direct
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]

