import { fetchGetFunct, fetchPostFunct } from "../fetch/fetch";

export interface Post {
  id?: string;
  title: string;
  content: string;
  user_id?: string;
  username?: string;
  created_at?: string; //timestamps can't be parsed directly into Date types in typescript.
  updated_at?: string; //you must call new Date(timestamp: string)
}

export async function fetchPostFeed(): Promise<Post[]> {
  const funct = fetchGetFunct<Post[]>("/posts");
  return await funct([]);

}

export async function fetchPost(id: string): Promise<Post | null> {
  const funct = fetchGetFunct<Post>(`/posts/${id}`);
  return await funct(null)
}

export async function createPost(post: Post): Promise<Error | null> {
  const func = fetchPostFunct<Post>("/posts");
  let res = await func(post);

  if (res.ok) {
    return null;
  }

  let content = await res.text();

  return new Error(content);
}
