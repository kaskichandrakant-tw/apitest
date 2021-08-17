FROM node:16-alpine

RUN mkdir -p /mountebannk

WORKDIR /mountebank

COPY package*.json ./

COPY . /mountebank

RUN npm install -g mountebank
