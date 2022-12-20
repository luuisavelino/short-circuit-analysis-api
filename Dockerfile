FROM golang:1.18-bullseye AS build-stage

WORKDIR /short-circuit-analysis-elements/

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

COPY ./main.go .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .


FROM alpine:3.17.0

WORKDIR /short-circuit-analysis-elements/

COPY --from=build-stage /short-circuit-analysis-elements/files/ ./files

COPY --from=build-stage /short-circuit-analysis-elements/main ./

EXPOSE 8080

CMD ["./main"]