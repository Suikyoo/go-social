
import { type RouteConfig } from "@mateothegreat/svelte5-router"; 

import Home from "@lib/pages/Home.svelte";
import NotFound from "@lib/pages/NotFound.svelte";
import PostCreation from "@lib/pages/PostCreation.svelte";
import PostView from "../pages/PostView.svelte";


//damn this router actually encourages using some compact regex with capture groups as params

//note: you can add a prehook to the routes, for authguards, redirecting the link to an "unauthorized" info page
//
//when doing an or in regex, you need to implement groups
//(?: ...) is there to specify a non-capturing group 
//Yes, you've initialized a group but you don't want it to,
//say in this case, for the value of the regex group to be part of the 
//route props, you don't need a member of the route prop to just say either "" or "home". 
//That data is useless
    
export const routes: RouteConfig[] = [
  {
    component: Home,
  },
  {
    component: NotFound,
    path: "*",
  },
  {
    component: Home,
    path: "/home/?",
  },
  {
    component: PostView,
    path: "/posts/(?<id>[0-9]{0,25})",
  },
  {
    component: PostCreation,
    path: "/posts/create",
  },

]

