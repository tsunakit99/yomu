import { getAllArticles } from "@/lib/posts/getAllArticles";
import { ArticleList } from "@/components/blog/ArticleList";

export default async function BlogPage() {
  const articles = await getAllArticles();

  return (
    <main className="p-8 max-w-3xl mx-auto">
      <h1 className="text-2xl font-bold mb-6">Yomu Blog</h1>
      <ArticleList articles={articles} />
    </main>
  );
}
