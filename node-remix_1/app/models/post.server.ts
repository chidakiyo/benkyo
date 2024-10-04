import { Post } from "@prisma/client";
import { prisma } from "~/db.server";

// interface Post {
//     slug: string;
//     title: string;
// }

export async function getPosts(): Promise<Post[]> {
    return prisma.post.findMany();
}

export async function getPost(slug: string): Promise<Post> {
    return prisma.post.findUnique({
        where: {
            slug,
        },
    });
}

export async function createPost(
    post: Pick<Post, "title" | "slug" | "markdown">
) {
    return prisma.post.create({ data: post });
}