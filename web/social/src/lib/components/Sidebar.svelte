<script>
  import SquareLibrary from "lucide-svelte/icons/square-library";
  import CirclePlus from "lucide-svelte/icons/circle-plus";
  import Search from "lucide-svelte/icons/search";
  import House from "lucide-svelte/icons/house";

</script>

<nav class="flex flex-row w-full lg:w-20 justify-evenly lg:justify-start lg:h-full lg:flex-col">
  <div class="p-5">
  <Search strokeWidth={1.75} />
  </div>

  <div class="p-5">
  <House strokeWidth={1.75} />
  </div>

    <div class="p-5">
  <SquareLibrary strokeWidth={1.75} />
  </div>

    <div class="p-5">
  <CirclePlus strokeWidth={1.75} />
  </div>

</nav>
