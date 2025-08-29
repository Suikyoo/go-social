import { apiUrl }  from "@lib/env/env"

export interface User {
  id?: string;
  username: string;
  password?: string;
}

export async function authenticateUser(u: User): Promise<Response>{

  try {
    return await fetch(apiUrl + "/auth/user", {
      method: "POST",
      headers: {
        'Accept': 'application/json',
        'Content-type': 'application/json',
      },
      body: JSON.stringify(u),
    });
  }
  catch (e) {
    throw new Error(e);
  }

}
