<script lang="ts">

  import Router from 'svelte-spa-router';
  import { SetAuthContext } from "@lib/auth/auth";
  import { routes } from "@lib/routes/route"

  import Navbar from "@components/Navbar.svelte"
  import Sidebar from "@components/Sidebar.svelte";


  //seed. next time, try to fetch for the auth in the login page and then create the context
  //SetAuthContext({id: BigInt(0), name: "test_user", src: "https://cdn.britannica.com/08/190808-050-CB26C47B/The-Powerpuff-Girls-Bubbles-Blossom-Buttercup.jpg"});

</script>

<main>

  <div>
    <Navbar/>
  </div>

  <div class="main_section">

    <Sidebar/>

    <div class="content">
        <Router {routes}/>
    </div>

  </div>

</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    width: 100vw;
    height: 100vh;
    overflow: hidden;
  }

  .main_section {
    display: flex;
    flex-direction: row;
    justify-content: start;
    flex: 1;
  }

  .content {
    width: 100%;
    border-top-left-radius: 3em;
    overflow: hidden;
  }

</style>
