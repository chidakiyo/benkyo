import { LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { SlowBuffer } from "buffer";
import invariant from "tiny-invariant";
import { getPost } from "~/models/post.server";
import { json } from "@remix-run/node";
import { marked } from "marked";

export const loader = async ({params}:LoaderFunctionArgs) => {
    const slug = params.slug;
    invariant(slug, "slug is requitred");

    const post = await getPost(slug);
    invariant(post, "Post not found.")

    const markdown = post.markdown;
    const html = marked(markdown)

    return json({post, html});
}

export default function PostSlut() {
    const {post, html} = useLoaderData<typeof loader>();
    return (
        <main className="mx-auto max-w-4x1">
            <h1 className="my-6 border-b-2 text-center text-3xl">{post.title}</h1>
            <div dangerouslySetInnerHTML={{__html: html}}></div>
        </main>
    );
}