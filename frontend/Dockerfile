FROM node:10.16.3-alpine

WORKDIR /app

RUN apk update && \
    npm install -Sg npm axios firebase

ADD package*.json ./

ADD . .

RUN npm install --no-optional

CMD ["npm", "run", "serve"]