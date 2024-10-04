import { json } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";
import { getPosts } from "~/models/post.server";
import { Outlet } from "@remix-run/react";

export const loader = async () =>  {
    return json({posts:await getPosts()})
}


export default function PostAdmin () {

    const {posts} = useLoaderData<typeof loader>();

    return (
        <div className="mx-auto max-w-4x1"> 
            <h1 className="text-center text-3xl my-6 mb-2 border-b-2">Blog Admin</h1>
            <div className="flex items-start gap-4">
                <nav>
                    <ul>{posts.map((post) => (
                        <li key={post.slug}>
                            <Link to={ "/posts/" + post.slug} className="text-blue-600 underline">{post.title}</Link>
                        </li>
                    ))}</ul>
                </nav>

                <main className="flex-grow">
                    <Outlet />
                </main>
            </div>
        </div>
    )
}