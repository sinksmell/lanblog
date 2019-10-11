FROM cargo.caicloud.xyz/library/debian:stretch

LABEL maintainer="sinksmell <cxxjcs@gmail.com>"

WORKDIR /root

COPY bin/lanblog /root/lanblog

COPY conf  /root/conf

ENTRYPOINT ["/root/lanblog"]

