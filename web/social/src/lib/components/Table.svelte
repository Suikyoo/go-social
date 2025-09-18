<script lang="ts">
  import { type Snippet } from "svelte";
  let {children, className}: {children: Snippet, className: string} = $props();


  $effect( () => {
    if (children) {
      let table_elements = document.querySelectorAll<HTMLDivElement>("#table > div");

      for (const [index, e] of table_elements.entries()) {
        e.style.boxSizing = "border-box";
        e.style.padding = "0.5em";

        if (index == 0) {
          continue;
        }

        e.style.borderTop = "inherit";
      }
    }
  });

</script>

<div id="table" class={className}>
  {@render children?.()}
</div>

<style>
  #table {
    display: grid;
    border-radius: 5px;
  }
</style>
