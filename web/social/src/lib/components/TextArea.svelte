<script lang="ts">
  let {className = "", name = "", value = $bindable(""), ...props} = $props();

</script>

<label for="textarea" class={`flex flex-col items-start bg-inherit ${className}`}>
  <p class="select-none">{name}</p>
  <!--className should be solely for label 
    since setting them both for parent and child 
    especially for relative sizing would mess things up
    -->
    <textarea {...props} {name} bind:value class={`rounded-sm border-[1px] border-slate-100 h-full w-full box-border px-4 py-2`} minlength=1 required></textarea>
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
  label > p:has(+ textarea:focus),
  label > p:has(+ textarea:valid){
    top: 0.7em;
    left: 0em;
    background-color: inherit;
    filter: brightness(1.0);
  }

</style>

