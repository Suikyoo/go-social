<script lang="ts">
  import { onMount } from "svelte";
  import { type Post, fetchPostFeed } from "@lib/types/post";

  let feed: Post[] = [];

  const maxContentLength = 10;

  onMount(async() => {
    feed = await fetchPostFeed();
  });

</script>

<section class="items-center w-full h-full overflow-y-scroll">

  <div class="flex w-[90vw] h-2/3 bg-slate-300 rounded-xl place-items-center">
      <h3 class="font-bold text-slate-950">Why scroll through reels when you can analyze taylor series expansions in <span class="text-red-800">this</span> bad boy?</h3>
  </div>

  <div class="flex flex-col justify-start p-5 mt-5 box-border">

    {#each feed as f}
      <div class="flex flex-col items-start justify-center w-full p-5 mb-5 rounded-xl bg-slate-100 box-border">
        <p class="text-left text-slate-800">{f.username}</p>
        <h2 class="font-bold text-slate-800">{f.title}</h2>
        <p class="text-slate-500">{f.content.length <= maxContentLength ? f.content : f.content.substring(0, maxContentLength) + '...'}</p>
                  
      </div>
    {/each}

  </div>

</section>

<style>
  .preface {
      color: var(--color-black);
    }

    .space {
      height: calc(100vh - 4em - 10px);
      overflow-y: scroll;
      box-sizing: border-box;
      padding: 2em;
      flex: 1;
    }

    .content {
      display: grid;
      margin-top: 3em;
      grid-template-columns: auto auto;
      row-gap: 1em;
      column-gap: 1em;
      box-sizing: border-box;
      padding: 1em;
      height: 100vh;
    }
</style>
