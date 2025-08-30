import Home from "@lib/pages/Home.svelte";
import NotFound from "@lib/pages/NotFound.svelte";
import PostCreation from "@lib/pages/PostCreation.svelte";

export const routes = {
    "/": Home, 
    "/posts/create": PostCreation,
    "*": NotFound,
};

