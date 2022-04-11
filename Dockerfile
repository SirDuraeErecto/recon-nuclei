FROM golang:1.17.6-alpine as build-env
RUN go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest

FROM alpine:3.15.4
RUN apk add --no-cache bind-tools ca-certificates chromium
COPY --from=build-env /go/bin/nuclei /usr/local/bin/nuclei
ENTRYPOINT ["nuclei"]
