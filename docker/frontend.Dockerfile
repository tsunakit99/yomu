# フロント: Next.js 15 + React 19 + pnpm
FROM node:20-alpine

WORKDIR /app

# pnpm
RUN corepack enable

# 依存ファイルコピー＆インストール
COPY frontend/pnpm-lock.yaml frontend/package.json ./
RUN pnpm install

# 残りのファイルコピー
COPY frontend ./

# ポート & 起動
EXPOSE 3000
CMD ["pnpm", "dev"]
