import { apiUrl }  from "../env/env"

//returns a postfetch function that takes in data of type T and forwards it to the server
//the postfetch function returned takes in data and returns a response
export function fetchPostFunct<T>(url: string): (data: T) => Promise<Response> {
  return async (data) => { 
    try {
      return await fetch(apiUrl + url, {
        method: "POST",
        headers: {
          'Accept': 'application/json',
          'Content-type': 'application/json',
        },
        body: JSON.stringify(data),
        credentials: "include",
      });
    }
    catch (e) {
      throw new Error(e);
    }
  }

  
}

//this getfetch function just takes nothing, fetches to the url, and then returns a data type of T
export function fetchGetFunct<T>(url: string): (baseData: T) => Promise<T> {
  return async (baseData) => {
    try {
      let res = await fetch(apiUrl + url, {
        method: "GET",
        credentials: "include",
      });
      return await res.json();

    }
    catch (e) {
      return baseData;
    }

  }
}
