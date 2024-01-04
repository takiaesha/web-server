#base image, main container
FROM golang:latest AS base

#not necessary to add
LABEL maintainer="https://github.com/takiaesha/web-server"

#creates a path in container named web-server
WORKDIR /web-server

#copy all files from current dir to conatiner dir
COPY . .

#container executes here and take this bianry file to multistage
RUN go build -o webserver

#multistage starts
FROM ubuntu:22.04

#creates a path in second container(root dir)
WORKDIR /web-server

#this image copies from base image, directory of base image and binary file-> /web-server/webserver
#and this container copies at light-server end within this directory
COPY --from=base /web-server/webserver /light-server

#port will flased at this port
EXPOSE 8080

#container starts config from here, sets default command
ENTRYPOINT [ "/light-server" ]






