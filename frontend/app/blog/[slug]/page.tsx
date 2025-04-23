import { getArticleBySlug } from "@/lib/posts/getArticleBySlug";
import { ArticleDetail } from "@/components/blog/ArticleDetail";
import { notFound } from "next/navigation";

export default async function BlogDetailPage({ params }: { params: { slug: string } }) {
  try {
    const article = await getArticleBySlug(params.slug);
    return (
      <main className="p-8 max-w-3xl mx-auto">
        <ArticleDetail article={article} />
      </main>
    );
  } catch {
    notFound();
  }
}