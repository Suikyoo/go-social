<script lang="ts">
  import { type Post, fetchPost } from "../types/post";


  let { route } = $props();
  let id = route.result.path.params.id;

</script>

<main class="size-full flex flex-col justify-start box-border p-5">
  {#await fetchPost(id)}
    <h1>Fetching post data...</h1>

  {:then post } 
    <section class="border-[0.5px] text-left rounded-sm">
      <div class="box-border p-2 flex flex-row justify-between">
        <h1 class="font-bold">{post?.title}</h1>
        <p>{post?.username}</p>
      </div>
      <hr />
      <div class="box-border p-5">
        <p>{post?.content}</p>
      </div>
    </section>

    <section class="">

    </section>

  {:catch err}
    <p>{err}</p>
  {/await}
  
</main>
