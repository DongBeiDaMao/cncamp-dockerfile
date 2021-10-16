FROM ubuntu
ENV VERSION=v1.0.0
ADD bin/amd64/my_httpserver /my_httpserver
ENTRYPOINT /my_httpserver