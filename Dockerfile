FROM golang:1.18-bullseye AS build-stage

WORKDIR /go/src/github.com/luuisavelino/short-circuit-analysis-elements/

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

COPY ./main.go .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .


FROM alpine:latest

COPY --from=build-stage /go/src/github.com/luuisavelino/short-circuit-analysis-elements/files/ ./files

COPY --from=build-stage /go/src/github.com/luuisavelino/short-circuit-analysis-elements/main ./run/

EXPOSE 8080

CMD ["./main"]