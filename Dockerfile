FROM golang:alpine as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -v -o gg-cli

FROM node:slim as terraform_installer
WORKDIR /tf
RUN apt-get update && apt-get upgrade -y
RUN apt-get install wget -y && apt-get install unzip -y
RUN wget https://releases.hashicorp.com/terraform/1.0.9/terraform_1.0.9_linux_amd64.zip
RUN unzip terraform_1.0.9_linux_amd64.zip && rm terraform_1.0.9_linux_amd64.zip


FROM node:lts-alpine3.14 as construct_builder
COPY --from=terraform_installer /tf/terraform /usr/bin/terraform
COPY --from=builder /app/gg-cli /usr/bin/gg-cli
RUN apk add --update --no-cache openssl curl ca-certificates jq
WORKDIR /app
RUN yarn global add cdktf-cli@0.7.0
RUN gg-cli typescript
RUN cdktf get
RUN gg-cli python
RUN cdktf get


FROM alpine:latest as prod
WORKDIR /app
COPY --from=terraform_installer /tf/terraform ./terraform
COPY --from=construct_builder /app/imports ./imports
COPY --from=construct_builder /app/.gen ./.gen
