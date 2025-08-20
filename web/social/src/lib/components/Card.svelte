<script>
  let {children, width="100%", height="100%", color_var="--color-white"} = $props();

</script>

<div class="preface" 
     style=
     "
      --preface-color: var({color_var}); 
      --preface-width: {width}; 
      --preface-height: {height};
      "
     >
     {@render children?.()}
</div>

<style>

  .preface {
    background-color: var(--preface-color);
    height: var(--preface-height);
    border-radius: 1em;
    display: flex;
    justify-content: center;
    align-items: center;
  }

</style>
