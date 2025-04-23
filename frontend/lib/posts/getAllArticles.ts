import fs from "fs/promises";
import path from "path";
import matter from "gray-matter";
import { ArticleMeta } from "./types";

export async function getAllArticles(): Promise<ArticleMeta[]> {
  const postsDir = path.join(process.cwd(), "posts");
  const files = await fs.readdir(postsDir);

  const articles = await Promise.all(
    files.map(async (filename) => {
      const filePath = path.join(postsDir, filename);
      const fileContents = await fs.readFile(filePath, "utf8");
      const { data } = matter(fileContents);
      return {
        slug: filename.replace(/\.md$/, ""),
        title: data.title,
        date: data.date,
        tags: data.tags || [],
      } as ArticleMeta;
    })
  );

  return articles.sort((a, b) => b.date.localeCompare(a.date));
}