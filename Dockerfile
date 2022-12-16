FROM golang:latest AS builder

WORKDIR /go/src/github.com/luuisavelino/short-circuit-analysis-elements/

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

COPY ./main.go .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .


FROM alpine:latest

WORKDIR /root/

COPY --from=0 /go/src/github.com/luuisavelino/short-circuit-analysis-elements/files/ ./files

COPY --from=0 /go/src/github.com/luuisavelino/short-circuit-analysis-elements/main ./

EXPOSE 8080

CMD ["./main"]