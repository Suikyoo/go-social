<script lang="ts">
  import Input from "../components/Input.svelte";
  import TextArea from "../components/TextArea.svelte";
</script>

<form class="flex flex-col items-center justify-stretch h-full w-full bg-inherit">
  <Input name="Title" type="text"/>
  <div class="h-200">
  <TextArea rows=20 name="Content" type="text" class="flex-2"/>
  </div>
</form>
