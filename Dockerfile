FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV PORT 8080
EXPOSE $PORT

COPY ./bin/crawler /
CMD ["/crawler"]