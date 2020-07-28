FROM node:alpine3.12 as build
RUN mkdir -p /build
WORKDIR /build
COPY . /build
RUN npm i --production

ENV PORT=3000
EXPOSE ${PORT}

CMD [ "npm", "start" ]
