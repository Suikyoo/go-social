<script lang="ts">
    import Card from "@components/Card.svelte";

  let feed: {title: string, color_var: string}[] = [
{title: "title1", color_var: "--color-rose-200"},
{title: "title2", color_var: "--color-cyan-200"},
{title: "title3", color_var: "--color-amber-200"},
  ]
</script>

<section class="bg-white h-[60%] overflow-y-scroll">

  <Card width="100%" height="60vh" color_var="--color-slate-100">
      <h3 class="preface">Why scroll through reels when you can analyze taylor series expansions in <span style="color: blue;">this</span> bad boy?</h3>
  </Card>

  <div class="content">

    {#each feed as f}
      <Card color_var={f.color_var}>
        <h1>{f.title}</h1>
      </Card>
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
