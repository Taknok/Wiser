### STAGE 1: Build ###

FROM node:carbon-alpine as builder

WORKDIR /hackdash

# Install needed libraries (curl, git g++, ...)
RUN apk add --update alpine-sdk

COPY package.json package-lock.json ./

RUN npm i

COPY . .

## Build the angular app in production mode and store the artifacts in dist folder
RUN npm run build:prod:aot


### STAGE 2: Setup ###

FROM nginx:1.13.3-alpine

## Copy our default nginx config
COPY nginx/default.conf /etc/nginx/conf.d/

## Remove default nginx website
RUN rm -rf /usr/share/nginx/html/*

## From 'builder' stage copy over the artifacts in dist folder to default nginx public folder
COPY --from=builder /hackdash/dist /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]


# docker build -t hackdash/app . 
# docker run -d -p 8080:80 hackdash/app