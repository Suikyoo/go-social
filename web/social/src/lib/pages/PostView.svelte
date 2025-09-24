<script lang="ts">
  import { Send } from "@lucide/svelte";
  import { fetchPost } from "../types/post";
  import { type Comment, fetchCommentFeed, createComment} from "../types/comment";
  import Input from "../components/Input.svelte";
  import Table from "@components/Table.svelte"
  import {onMount} from "svelte";

  let { route } = $props();
  let id: string = route.result.path.params.id;

  let commentFeed: Comment[] = $state([]);

  let comment: Comment =  $state({post_id: id, content: ""});

  async function create() {
    let err = await createComment(comment);
    if (!err) {
      commentFeed.push({...comment})
      comment.content = "";
    }
  }

  onMount(() => {
    fetchCommentFeed(id).then((feed) => commentFeed = feed)
  })

</script>

<main class="size-full flex flex-col justify-start box-border p-5 border-zinc-800">
  {#await fetchPost(id)}
    <h1>Fetching post data...</h1>

  {:then post} 
    <Table className="grid-rows-1 text-left border-[0.5px] border-zinc-800 mb-10">
      <div class="flex flex-row justify-between">
        <h1 class="font-bold">{post.title}</h1>
        <div class="flex flex-row text-sm">
          <p>{post.username}</p>
          |
          <p>{new Date(post.created_at).toLocaleString()}</p>

        </div>
      </div>
      <div class="h-50">{post.content}</div>
    </Table>

    <Table className="grid-rows-1 text-left border-[0.5px] border-zinc-800"> 
      <div>Comments</div>
      {#each commentFeed as comment}
        <div class="flex flex-col w-full border-red-200">
          {comment.username}: {comment.content}
        </div>
      {/each}
      <div class="flex flex-row justify-between items-center w-full">
        <Input type="text" bind:value={comment.content} className="h-[2em] w-full" nativeClass="border-zinc-600 border-1"/>
        <button onclick={create} class="hover:outline-1 m-2 p-1 rounded-sm">
          <Send />
        </button>

      </div>

    </Table>

  {:catch err}
    <p>{err}</p>
  {/await}
  
</main>
