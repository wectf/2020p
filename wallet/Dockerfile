FROM node:14.2.0
WORKDIR /home/wallet/
COPY app.js ./app.js
COPY index.js ./index.js
COPY package.json ./package.json
COPY views ./views
COPY bin ./bin
RUN npm i
CMD ["bin/www"]