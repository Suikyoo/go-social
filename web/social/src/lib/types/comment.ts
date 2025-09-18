import { fetchGetFunct, fetchPostFunct } from "../fetch/fetch";

export interface Comment {
  id?: string;
  user_id?: string;
  post_id: string;
  username?: string;
  content: string;
  created_at?: string;
  updated_at?: string;
  likes?: string;
}


export async function fetchCommentFeed(postID: string): Promise<Comment[]> {
  const funct = fetchGetFunct<Comment[]>(`/comments/${postID}`);
  return await funct([]);
}

export async function createComment(comment: Comment): Promise<Error | null> {
  const funct = fetchPostFunct<Comment>(`/comments`);
  let res = await funct(comment);

  if (res.ok) {
    return null;
  }

  let content = await res.text();

  return new Error(content);

}
