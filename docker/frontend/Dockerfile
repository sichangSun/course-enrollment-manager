FROM node:22-alpine3.18

WORKDIR /app

COPY ./frontend/ ./

RUN apk update

CMD ["npm","run","dev","--host"]
