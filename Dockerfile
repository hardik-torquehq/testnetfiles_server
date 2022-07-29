FROM golang:1.18.4-alpine3.16
WORKDIR /app
COPY . .
RUN go build -o testnetfiles testnetfiles_server.go 




EXPOSE 8081 

CMD ["go", "run", "." ] 
