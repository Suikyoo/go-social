
import { fetchPostFunct } from "@lib/fetch/fetch"; 
export interface User {
  id?: string;
  username: string;
  password?: string;
  src?: string;
}


export async function authenticateUser(user: User): Promise<Error> | null {
  const func = fetchPostFunct<User>("/auth/token");
  let res = await func(user);

  if (res.ok) {
    return null;
  }
  let content = await res.text();

  return new Error(content);
}

export async function createUser(user: User): Promise<Error> | null {
  const func = fetchPostFunct<User>("/auth/user");
  let res = await func(user);

  if (res.ok) {
    return null;
  }

  let content = await res.text();

  return new Error(content);

}
