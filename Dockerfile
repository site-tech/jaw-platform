FROM golang:1.19.3-alpine3.16 AS build
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN apk add --update make
RUN make

FROM alpine:3.16
EXPOSE 8880
COPY --from=build /app/bin/subgraph-temp /subgraph-temp
COPY .env .env
CMD [ "/subgraph-temp"]
