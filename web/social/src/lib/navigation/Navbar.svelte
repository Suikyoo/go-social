<script lang="ts">
  import Profile from "./Profile.svelte";
</script>


<nav>
  <div class="title">
    <h2>solu<span>Shown</span></h2>
  </div>

  <div class="profile">
    <Profile/>
  </div>

</nav>

<style>

  nav {
    width: 100%;
    height: 4em;
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

  .title {
    padding-left: 20px;
    font-weight: 200;

  }

  .profile {
    display: flex;
    padding-right: 20px;
  }


</style>
