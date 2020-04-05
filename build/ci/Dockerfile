FROM golang:1.12-alpine
ENV CGO_ENABLED=0 GO111MODULE=on
RUN apk add --no-cache git
ADD https://raw.githubusercontent.com/eficode/wait-for/master/wait-for /usr/bin/wait-for
RUN chmod +x /usr/bin/wait-for
RUN go get golang.org/x/lint/golint
CMD /usr/bin/wait-for segmenterv1:8080 -t 60 -- go test github.com/ndthuan/go-vi-wordseg-client/...
