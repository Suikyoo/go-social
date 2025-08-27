import { apiUrl }  from "@lib/env/env"

export interface Post {
  title: string;
  content: string;
  username: string;
  user_id: string;
  create_at: Date;
  updated_at: Date;
}

export async function fetchPostFeed (): Promise<Post[]> {
  try {
    console.log("hi")
    console.log(apiUrl)
    let res = await fetch(apiUrl + "/posts");
    console.log(res)
    if (!res.ok) {
      let text = await res.text();
      throw new Error(`${res.status}: ${text}`);

    }
    return await res.json();
  }
  catch (e) {
    console.log(e);
    return [];
  }

}

