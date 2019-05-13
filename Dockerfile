FROM golang:alpine as source
WORKDIR /home/server
COPY . .
WORKDIR cmd/crypto-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -mod vendor -o crypto-service

FROM alpine as runner
LABEL name="bitstored/crypto-service"
RUN apk --update add ca-certificates
COPY --from=source /home/server/cmd/crypto-service/crypto-service /home/crypto-service
COPY --from=source /home/server/scripts/localhost.* /home/scripts/
WORKDIR /home
EXPOSE 4004
EXPOSE 5004
ENTRYPOINT [ "./crypto-service" ]
