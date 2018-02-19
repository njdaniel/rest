FROM alpine:3.7
ADD rest /usr/bin/rest
ENTRYPOINT ["rest"]