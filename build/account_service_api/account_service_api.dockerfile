FROM alpine:latest

RUN mkdir /app

COPY ./account_service_api /app

EXPOSE 8080

# CMD ["/app/account_service_api"]
ENTRYPOINT [ "/app/account_service_api" ]