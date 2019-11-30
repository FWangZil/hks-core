FROM golang:stretch as golang-builder
WORKDIR /go/src/code.aliyun.com/yjqu/apigate-card
COPY . .
COPY cmd/card/main.go ./main.go
#RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=golang-builder /go/src/code.aliyun.com/yjqu/apigate-card/app .
COPY --from=golang-builder /go/src/code.aliyun.com/yjqu/apigate-card/docs ./docs
COPY --from=golang-builder /go/src/code.aliyun.com/yjqu/apigate-card/template ./template
EXPOSE 8080
ENTRYPOINT ["./app"]