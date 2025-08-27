<script lang="ts">
  import Profile from "./Profile.svelte";
</script>


<nav class="flex flex-row justify-between w-full px-5 h-18 box-border">
  <div class="flex flex-row items-center">
    <h2>I<span>dunno</span></h2>
    <input class="px-3 mx-10 text-sm bg-transparent min-w-20 rounded-xs text-zinc-100 box-border" type="text">
  </div>

  <div class="flex place-items-center">
    <Profile/>
  </div>

</nav>
