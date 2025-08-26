<script lang="ts">

  import Router from 'svelte-spa-router';
  import { routes } from "@lib/routes/route"

  import Navbar from "@components/Navbar.svelte"
  import Sidebar from "@components/Sidebar.svelte";


  //seed. next time, try to fetch for the auth in the login page and then create the context
  //SetAuthContext({id: BigInt(0), name: "test_user", src: "https://cdn.britannica.com/08/190808-050-CB26C47B/The-Powerpuff-Girls-Bubbles-Blossom-Buttercup.jpg"});

</script>

<main class="flex flex-col w-screen h-screen overflow-hidden">

  <div>
    <Navbar/>
  </div>

  <div class="flex flex-row justify-start flex-1">

    <Sidebar/>

    <div class="w-full rounded-tl-xl overflow-hidden">
        <Router {routes}/>
    </div>

  </div>

</main>
