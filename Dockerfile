# build stage
FROM golang:1.19-alpine AS builder

LABEL maintainer="Joel Santos <joe@joesantos.io>"

WORKDIR /app
COPY . .

RUN apk update && apk add --virtual build-dependencies build-base gcc git make

RUN go mod download
RUN go build -v -o ./cmd

# build public stage
FROM node:18-alpine AS publicbuilder

ENV TZ="Etc/UTC"
ENV ENV="production"
ENV NODE_ENV="production"

WORKDIR /app
COPY ./app .

RUN apk --no-cache add --virtual native-deps python3 g++ gcc libgcc libstdc++ linux-headers make git

RUN NODE_ENV=development npm install --production=false
RUN npm run build:production

# final stage
FROM alpine

ARG TZ="Etc/UTC"
ARG ENV="production"
ARG REDIS_HOST="localhost:6379"
ARG REDIS_SECRET="redis"

# use built variables as defaults for the ENV
ENV TZ=${TZ}
ENV ENV=${ENV}
ENV REDIS_HOST=${REDIS_HOST}
ENV REDIS_SECRET=${REDIS_SECRET}

ENV PORT=4040

WORKDIR /app
COPY --from=builder /app/cmd ./cmd

RUN mkdir -p ./app
COPY --from=publicbuilder /app/dist ./app/dist

CMD ["./cmd"]
