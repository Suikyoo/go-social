<script lang="ts">
  let {name = "", value = $bindable(""), ...props} = $props();

</script>

<label class="flex flex-col items-start bg-inherit">
  <p class="select-none">{name}</p>
  <input {...props} {name} bind:value class="rounded-sm border-[1px] border-slate-100 h-[2em] box-border px-5 py-4" minlength=1 required/>
</label>

<style>

  label > p {
    text-align: left;
    position: relative;
    top: 1.6em;
    left: 1em;
    margin: 0;
    filter: brightness(0.4);
    background-color: transparent;
    transition: top 0.2s, left 0.2s, filter 0.2s, background-color 0.2s;

  }
  label > p:has(+ input:focus),
  label > p:has(+ input:valid){
    top: 0.7em;
    left: 0em;
    background-color: inherit;
    filter: brightness(1.0);
  }

</style>

