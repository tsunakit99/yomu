# フロント: Next.js 15 + React 19 + pnpm
FROM node:20

WORKDIR /app

# pnpm の使用を有効化
RUN corepack enable

# 必要ファイルのみコピーして依存インストール
COPY package.json pnpm-lock.yaml ./
RUN pnpm install

# 残りのアプリファイルをコピー（node_modules を上書きしないようにこの順序）
COPY . .

EXPOSE 3000
CMD ["pnpm", "dev"]
