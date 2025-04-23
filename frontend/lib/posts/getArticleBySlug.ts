import fs from "fs/promises";
import path from "path";
import matter from "gray-matter";
import { ArticleMeta } from "./types";
import { remark } from "remark";
import html from "remark-html";
import remarkGfm from "remark-gfm";

export type ArticleContent = ArticleMeta & { contentHtml: string };

export async function getArticleBySlug(slug: string): Promise<ArticleContent> {
  const filePath = path.join(process.cwd(), "posts", `${slug}.md`);
  const fileContent = await fs.readFile(filePath, "utf8");
  const { data, content } = matter(fileContent);

  const processedContent = await remark().use(remarkGfm).use(html).process(content);
  const contentHtml = processedContent.toString();

  return {
    slug,
    title: data.title,
    date: data.date,
    tags: data.tags || [],
    contentHtml,
  };
}