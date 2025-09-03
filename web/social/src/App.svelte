<script lang="ts">

  import Router from 'svelte-spa-router';
  import { routes } from "@lib/routes/route"

  import Navbar from "@components/Navbar.svelte"
  import Sidebar from "@components/Sidebar.svelte";
  import Login from './lib/components/Login.svelte';
  import Signup from './lib/components/Signup.svelte';


  //seed. next time, try to fetch for the auth in the login page and then create the context
  //SetAuthContext({id: BigInt(0), name: "test_user", src: "https://cdn.britannica.com/08/190808-050-CB26C47B/The-Powerpuff-Girls-Bubbles-Blossom-Buttercup.jpg"});

</script>

<main class="flex flex-col w-screen h-screen justify-start overflow-hidden">

  <div>
    <Navbar/>
  </div>

  <div class="flex w-screen flex-col-reverse flex-1 overflow-y-hidden lg:flex-row">

    <Login/>
    <Signup/>
    <Sidebar/>
    <div class="w-full h-full overflow-hidden rounded-tl-xl bg-zinc-950">
      <Router { routes }></Router>
    </div>

  </div>

</main>
