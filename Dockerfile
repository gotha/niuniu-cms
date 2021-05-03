FROM golang:1

COPY . /
WORKDIR /

RUN CGO_ENABLED=0 go build -mod=readonly -a -o /artifacts/niuniu-cms

FROM scratch
WORKDIR /
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /artifacts/* /

CMD [ "/niuniu-cms" ]

