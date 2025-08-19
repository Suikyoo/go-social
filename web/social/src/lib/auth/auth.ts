import { getContext, setContext } from "svelte";
import { writable, type Writable } from "svelte/store";

const AuthContext = Symbol("auth");

export interface User {
  id: BigInt;
  name: string;
  src?: string;
};



//this one only needs the user object
export function SetAuthContext(user: User) {
  setContext(AuthContext, writable(user));
}

//this one gets the store object
export function GetAuthContext(): Writable<User> {
  return getContext<Writable<User>>(AuthContext);
}

