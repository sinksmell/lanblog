FROM node:14.16.1-alpine3.13 AS JS_BUILD
COPY webapp /webapp
WORKDIR /webapp/app
RUN npm install && npm run build

FROM golang:1.16.3-alpine3.13 AS GO_BUILD
RUN apk update && apk add build-base
COPY server /server
WORKDIR /server
RUN GOPROXY=https://goproxy.cn go mod tidy
RUN go build -o /go/bin/server

FROM alpine:3.13.5
COPY --from=JS_BUILD /webapp/app/build* ./webapp/app
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
