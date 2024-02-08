FROM golang:1.21.2

ARG bell_hostname=127.0.0.1
ARG bell_port=8099
ARG bell_email=
ARG bell_email_pwd=

ENV BELL_HOSTNAME=${bell_hostname}
ENV BELL_PORT=${bell_port}
ENV BELL_EMAIL=${bell_email}
ENV BELL_EMAIL_PWD=${bell_email_pwd}

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

EXPOSE ${bell_port}

CMD ["bell"]