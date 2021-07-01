FROM golang:alpine as retina
RUN apk --update add ca-certificates gcc libc-dev
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o retina .

FROM node:16 AS react
WORKDIR /usr/src/app
COPY ./web/app/package.json ./
RUN npm install
COPY ./web/app/ ./
RUN npm run build

FROM alpine:latest
COPY --from=retina /app/retina .
COPY --from=retina /app/config.yaml /config.yaml
COPY --from=retina /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=react /usr/src/app/build/ /web/app/build
CMD ["/retina"]
