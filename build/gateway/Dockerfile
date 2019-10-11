FROM nginx:1.17

LABEL maintainer="sinksmell <cxxjcs@gmail.com>"

WORKDIR /root

COPY front /root/front

COPY front/nginx.conf  /etc/nginx/nginx.conf

COPY front/conf  /etc/nginx/conf.d

