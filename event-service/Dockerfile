FROM golang:1.22.0
ENV APP_NAME=event-service
ENV APP_PATH=/go/$APP_NAME
RUN mkdir -p $APP_PATH
COPY . $APP_PATH
WORKDIR . $APP_PATH
RUN go build -o $APP_NAME .
EXPOSE 8080
CMD ["./event-service"]