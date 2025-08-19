<script lang="ts">
  import Navbar from "./lib/navigation/Navbar.svelte";

  import { SetAuthContext } from "./lib/auth/auth";
  import Sidebar from "./lib/navigation/Sidebar.svelte";
  import Home from "./lib/home/Home.svelte";


  //seed. next time, try to fetch for the auth in the login page and then create the context
  SetAuthContext({id: BigInt(0), name: "test_user"});

</script>

<main>
  <div>
    <Navbar/>
  </div>

  <div class="main_section">
    <Sidebar/>

    <div class="content">
      <Home/>
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
