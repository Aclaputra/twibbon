# stage 1
FROM node:20-alpine AS builder


WORKDIR /usr/src/app


# COPY package*.json ./
COPY ./ /usr/src/app/


RUN npm install

RUN npm run build

# COPY . .


# stage 2
RUN npm run build --prod


FROM nginx:latest


RUN rm -rf /usr/share/nginx/html/*


COPY --from=builder /usr/src/app/dist/twibbon-angular /usr/share/nginx/html


COPY nginx.conf /etc/nginx/nginx.conf


EXPOSE 80
# EXPOSE 443
