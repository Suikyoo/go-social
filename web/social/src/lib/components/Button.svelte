<script lang="ts">
  let {className = "", children, ...props} = $props();

</script>

<button {...props} class={`px-4 py-2 hover:outline transition:outline ${className}`}>{@render children?.()}</button>

<style>
</style>
