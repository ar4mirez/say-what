FROM php:7.4-cli
COPY . /usr/src/app
WORKDIR /usr/src/app

ENV PORT=3000
ENV APP_URL=0.0.0.0:${PORT}

EXPOSE ${PORT}

RUN chmod +x ./entrypoint.sh

ENTRYPOINT [ "./entrypoint.sh" ]
