import { ActionFunctionArgs, redirect } from "@remix-run/node"
import { Form, useNavigation } from "@remix-run/react"
import { createPost } from "~/models/post.server";

export const action = async ({request}: ActionFunctionArgs) => {

    await new Promise((resolve) => setTimeout(resolve, 1000))

    const formData = await request.formData();

    const title = formData.get("title") as string;
    const slug = formData.get("slug") as string;
    const markdown = formData.get("markdown") as string;

    await createPost({title, slug, markdown})

    return redirect("/posts/admin")
}

export default function NewPost() {

    const navigation = useNavigation();
    const isCreating = Boolean(navigation.state === "submitting")

    return (
        <Form method="post" className="space-y-3" >
            <div>
                <label htmlFor="title">Post Title:</label>
                <input type="text" name="title" className="w-full rounded border border-gray-500 px-2 py-1 text-lg" />
            </div>
            <div>
                <label htmlFor="title">Post Slug:</label>
                <input type="text" name="slug" className="w-full rounded border border-gray-500 px-2 py-1 text-lg" />
            </div>
            <div>
                <label htmlFor="markdown">Post Slug:</label>
                <br />
                <textarea id="markdown" name="markdown"
                rows={20}
                className="w-full rounded border border-gray-500 px-2 py-1 text-lg" />
            </div>

            <div className="text-right">
                <button 
                type="submit"
                className="bg-blue-500 py-2 px-4 text-white">
                    {isCreating ? "Submitting..." : "Create Post"}
                </button>
            </div>
        </Form>
    )
}