FROM node:18.13.0-alpine

WORKDIR /usr/src/app

COPY package*.json ./

RUN yarn

EXPOSE 3000

CMD ["yarn", "dev"]