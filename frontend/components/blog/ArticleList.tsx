import { ArticleMeta } from "@/lib/posts/types";
import { ArticleCard } from "./ArticleCard";

export function ArticleList({ articles }: { articles: ArticleMeta[] }) {
  return (
    <div className="grid gap-6">
      {articles.map((article) => (
        <ArticleCard key={article.slug} article={article} />
      ))}
    </div>
  );
}
