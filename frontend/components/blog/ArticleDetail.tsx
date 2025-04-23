import { ArticleContent } from "@/lib/posts/getArticleBySlug";
import { Badge } from "@/components/ui/badge";

export function ArticleDetail({ article }: { article: ArticleContent }) {
  return (
    <article className="prose dark:prose-invert max-w-none">
      <h1>{article.title}</h1>
      <p className="text-sm text-gray-500">{article.date}</p>
      <div className="flex gap-2 mb-4">
        {article.tags.map((tag) => (
          <Badge key={tag} variant="secondary">{tag}</Badge>
        ))}
      </div>
      <div dangerouslySetInnerHTML={{ __html: article.contentHtml }} />
    </article>
  );
}