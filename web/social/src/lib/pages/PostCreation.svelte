<script lang="ts">
  import Input from "../components/Input.svelte";
  import TextArea from "../components/TextArea.svelte";
  import Button from "../components/Button.svelte";
  import { type Post, createPost } from "../types/post"; 

  import { pop } from "@mateothegreat/svelte5-router";

  let post: Post = $state({title: "", content: ""});
  let feedback: HTMLParagraphElement | null = $state(null);
  
  async function create() {
    let err = await createPost(post);
    if (err){
      feedback.innerText = err.message;
      feedback.style.color = "red";
    }
    else {
      feedback.innerText = "post created successfully";
      feedback.style.color = "green";
      pop();
    }

  }

</script>

<form onsubmit={(e) => {e.preventDefault(); create()}} class="flex flex-col items-center justify-start h-full w-full bg-inherit overflow-y-scroll box-border p-10">
  <div class="flex flex-col items-center w-full max-w-[640px] bg-inherit">
    <h1 class="self-start">Create Post</h1>
    <Input name="Title" type="text" className="w-full h-[4em]" bind:value={post.title} />
    <TextArea rows=20 name="Content" type="text" className="w-full h-[10em]" bind:value={post.content}/>
    <p bind:this={feedback}></p>
    <Button className="bg-blue-700 focus-visible:outline-blue-700 hover:outline-blue-700 self-end">Submit</Button>
  </div>

</form>
