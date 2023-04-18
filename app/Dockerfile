FROM golang:alpine

RUN mkdir -p /var/www

WORKDIR /var/www

COPY . /var/www
# change GOARCH for your os, yours will probably be amd64
# there's a bug with qemu on the M1 chip, so I need arm
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o app

EXPOSE 3200
ENTRYPOINT ["/var/www/app"]