<script>
  import SquareLibrary from "lucide-svelte/icons/square-library";
  import CirclePlus from "lucide-svelte/icons/circle-plus";
  import Search from "lucide-svelte/icons/search";
  import House from "lucide-svelte/icons/house";

</script>

<nav>
  <div>
  <Search strokeWidth={1.75} />
  </div>

  <div>
  <House strokeWidth={1.75} />
  </div>

  <div>
  <SquareLibrary strokeWidth={1.75} />
  </div>

  <div>
  <CirclePlus strokeWidth={1.75} />
  </div>

</nav>

<style>
  nav {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: start;
  }

  nav * {
    padding: 15px;
    padding-right: 2em;
  }
</style>
