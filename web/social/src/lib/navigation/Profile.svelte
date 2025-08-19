
<script lang="ts">
  import { GetAuthContext } from "../auth/auth"
  import CircleUser from "lucide-svelte/icons/circle-user";

  const user = GetAuthContext(); 

</script>

{#if $user.src !== undefined} 
<img src={$user.src} alt="user {$user.name}"/>
{:else}
  <CircleUser size=32/>
{/if}

<style>

  img {
    border-radius: 100px;
    border: 2px solid white;
  }

</style>

  




