FROM golang:1.22.0
ENV APP_NAME=user-service
ENV APP_PATH=/go/src/$APP_NAME
RUN mkdir -p $APP_PATH
COPY . $APP_PATH
WORKDIR $APP_PATH
RUN go build -o $APP_NAME
CMD ["./user-service"]