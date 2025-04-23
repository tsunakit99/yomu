FROM node:20

WORKDIR /app

# パッケージファイルだけ先にコピーしてインストールキャッシュを活用
COPY package.json package-lock.json* ./
RUN npm install

# 残りのソースをコピー
COPY . .

CMD ["npm", "run", "dev"]