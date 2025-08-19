<script lang="ts">
  import Preface from "./Preface.svelte";

  let feed: {title: string, color_var: string}[] = [
{title: "title1", color_var: "--color-rose-200"},
{title: "title2", color_var: "--color-cyan-200"},
{title: "title3", color_var: "--color-amber-200"},
  ]
</script>
<div class="space">

  <Preface width="100%" height="60vh" color_var="--color-slate-100">
    <h1>So...</h1>
  </Preface>

  <div class="content">

    {#each feed as f}
      <Preface color_var={f.color_var}>
        <h1>{f.title}</h1>
      </Preface>
    {/each}

  </div>

</div>

<style>
  h1 {
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
