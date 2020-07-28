FROM node:alpine3.12

RUN mkdir -p /app
WORKDIR /app

COPY . /app

RUN npm i

ENV PORT=3000
EXPOSE ${PORT}

CMD [ "npm", "start" ]
