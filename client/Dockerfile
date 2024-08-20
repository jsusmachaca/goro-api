FROM node:20-alpine3.20

WORKDIR /app

COPY . .

RUN npm i
RUN npm run build

# CMD [ "npm", "run", "dev" ]
CMD [ "node", "dist/index.js" ]