import { apiUrl }  from "@lib/env/env"

export interface User {
  id?: string;
  username: string;
  password?: string;
  src?: string;
}

function postFunct(url: string): (user: User) => Promise<Response> {
  return async (user: User) => { 
    try {
      return await fetch(apiUrl + url, {
        method: "POST",
        headers: {
          'Accept': 'application/json',
          'Content-type': 'application/json',
        },
        body: JSON.stringify(user),
      });
    }
    catch (e) {
      throw new Error(e);
    }
  }

  
}

export const authenticateUser = postFunct("/auth/token");
export const createUser = postFunct("/auth/user");
