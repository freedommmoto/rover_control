#build stage
FROM golang:1.18-alpine3.15 as builder
WORKDIR /app
COPY . .
RUN go get -u ./...
RUN cd /app/cmd && go build main.go

#last stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/cmd/main /app/cmd/main
ADD /instructions_file/instructions.txt /app/instructions_file/instructions.txt

EXPOSE 7471
CMD ["/app/cmd/main"]


