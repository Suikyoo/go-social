import { apiUrl }  from "@lib/env/env"

export interface Post {
  title: string;
  content: string;
  username: string;
  user_id: string;
  created_at: string; //timestamps can't be parsed directly into Date types in typescript.
  updated_at: string; //you must call new Date(timestamp: string)
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

