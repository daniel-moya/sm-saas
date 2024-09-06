FROM oven/bun

WORKDIR /app

COPY package.json .

RUN bun install

COPY . .

EXPOSE 50053

CMD [ "bun", "run", "index.ts" ]
