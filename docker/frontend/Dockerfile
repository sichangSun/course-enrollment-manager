FROM node:22-alpine3.18

WORKDIR /app

COPY ./frontend/package*.json ./

RUN npm install

COPY ./frontend/ ./

RUN apk update

CMD ["npm","run","dev","--host"]
