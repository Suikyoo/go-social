<script lang="ts">
  import { type Post, fetchPostFeed } from "@lib/types/post";
  import { userProfile } from "../components/Profile.svelte";
  import { goto } from "@mateothegreat/svelte5-router";

  let feed: Post[] = $state([]);

  const maxContentLength = 10;
  
  $effect(() => {
    if (userProfile.username) {
      //TECHNICALLY, this is superduper correct way of using effects. -_-
      //feed wont infinitely update as the variable is asynchronously scoped
      fetchPostFeed().then((value) => {feed = value});
    }
  });

</script>

<main class="flex flex-col items-center w-full h-full overflow-y-scroll bg-inherit">

  <div class="flex w-[90vw] h-2/3 bg-zinc-300 rounded-xl items-center justify-center box-border p-10">
      <h3 class="font-bold text-zinc-950">Why scroll through reels when you can analyze taylor series expansions in <span class="text-red-800">this</span> bad boy?</h3>
  </div>

  <div class="flex flex-col w-full justify-start p-5 mt-5 box-border">

    {#each feed as f}
      <button onclick={() => {goto(`/posts/${f.id}`)}} class="flex flex-col items-start justify-center w-full p-5 mb-5 rounded-xl bg-zinc-100 box-border">
        <p class="text-left text-zinc-800">{f.username}</p>
        <h2 class="font-bold text-zinc-800">{f.title}</h2>
        <div class="flex flex-row justify-between w-full">
          <p class="text-zinc-500">{f.content.length <= maxContentLength ? f.content : f.content.substring(0, maxContentLength) + '...'}</p>
          <p class="text-zinc-400 text-right">{f.created_at}</p>

        </div>

      </button>
    {/each}

  </div>

</main>
