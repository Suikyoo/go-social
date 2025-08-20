import Home from "@lib/pages/Home.svelte";
import NotFound from "@lib/pages/NotFound.svelte";

export const routes = {
    '/': Home, 
    '*': NotFound,
};

