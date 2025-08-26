<script lang="ts">
  import Profile from "./Profile.svelte";
</script>


<nav>
  <div class="left">
    <h2>I<span>dunno</span></h2>
    <input class="search" type="text">
  </div>

  <div class="right">
    <Profile/>
  </div>

</nav>

<style>
nav {
    width: 100%;
    height: 4.3em;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
    box-sizing: border-box;
  }

  span {
    color: var(--color-yellow-200);
    font-weight: 900;
    font-style: italic;
  }

  .left {
    margin-left: 20px;
    font-weight: 200;
    display: flex;
    flex-direction: col;
    justify-content: space-between;
    align-items: center;

  }

  .right {
    display: flex;
    padding-right: 20px;
  }

  .search {
    width: 8em;
    height: 1.5em;
    color: black;
    box-sizing: border-box;
    padding-left: 1em;
    padding-right: 1em;
    
  }


</style>
