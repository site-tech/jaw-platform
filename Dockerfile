FROM golang:1.19.3-alpine3.16 AS build
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN apk add --update make
RUN go mod download
RUN make

FROM alpine:3.16
EXPOSE 8880
COPY --from=build /app/bin/jaw-platform /jaw-platform
COPY .env .env
CMD [ "/jaw-platform"]
