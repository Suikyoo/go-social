<script lang="ts">

  import Router from 'svelte-spa-router';
  import { routes } from "@lib/routes/route"

  import Navbar from "@components/Navbar.svelte"
  import Sidebar from "@components/Sidebar.svelte";


  //seed. next time, try to fetch for the auth in the login page and then create the context
  //SetAuthContext({id: BigInt(0), name: "test_user", src: "https://cdn.britannica.com/08/190808-050-CB26C47B/The-Powerpuff-Girls-Bubbles-Blossom-Buttercup.jpg"});

</script>

<main class="flex flex-col w-screen h-screen justify-start">

  <div>
    <Navbar/>
  </div>

  <div class="flex w-screen flex-col-reverse flex-1 overflow-hidden lg:flex-row">

    <Sidebar/>
    <div class="w-full h-full overflow-hidden rounded-tl-xl">
      <Router {routes}/>
    </div>

  </div>

</main>
