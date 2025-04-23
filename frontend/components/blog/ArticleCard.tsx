import { ArticleMeta } from "@/lib/posts/types";
import Link from "next/link";

export function ArticleCard({ article }: { article: ArticleMeta }) {
  return (
    <div className="border p-4 rounded-lg shadow-sm">
      <h2 className="text-xl font-bold mb-1">
        <Link href={`/blog/${article.slug}`}>{article.title}</Link>
      </h2>
      <p className="text-sm text-gray-500">{article.date}</p>
      <div className="mt-2 text-xs text-slate-600 space-x-2">
        {article.tags.map((tag) => (
          <span key={tag} className="px-2 py-1 bg-slate-100 rounded">
            {tag}
          </span>
        ))}
      </div>
    </div>
  );
}