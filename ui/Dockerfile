FROM node:14-alpine AS development
WORKDIR /app

COPY ./package*.json ./
RUN npm install

EXPOSE 8081 9876

ENV PATH $PATH:/app/node_modules/.bin
CMD npm run dev


FROM node:14-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
RUN npm audit fix
COPY . .
# RUN npm run lint
RUN npm run build


FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/public /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]


FROM node:14-alpine as staging-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
RUN npm audit fix
COPY . .
# RUN npm run lint
RUN npm run build staging


FROM nginx:stable-alpine as production-staging
COPY --from=staging-stage /app/public /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]