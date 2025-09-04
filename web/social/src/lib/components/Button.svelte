<script lang="ts">
  let {className = "", children, ...props} = $props();

</script>

<button {...props} class={`px-4 py-2 my-3 rounded-sm outline-transparent outline-3 outline-offset-3 transition ${className}`}>{@render children?.()}</button>

