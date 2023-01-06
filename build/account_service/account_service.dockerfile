FROM alpine:latest

RUN mkdir /app

COPY ./account_service /app

EXPOSE 5055

ENTRYPOINT [ "/app/account_service" ]