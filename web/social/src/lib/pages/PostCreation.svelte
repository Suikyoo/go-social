<script lang="ts">
  import Input from "../components/Input.svelte";
  import TextArea from "../components/TextArea.svelte";
  import Button from "../components/Button.svelte";
</script>

<form class="flex flex-col items-start justify-start h-full w-full bg-inherit overflow-y-scroll box-border px-10">
  <h1 class="self-start">Create Post</h1>
  <Input name="Title" type="text" className="w-[90%] h-[4em]"/>
  <TextArea rows=20 name="Content" type="text" className="w-[90%] h-[10em]"/>
  <Button>Submit</Button>

</form>
