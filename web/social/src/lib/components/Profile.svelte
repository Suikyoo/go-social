<script lang="ts">
  import { SetAuthContext, GetAuthContext } from "../auth/auth"
  import CircleUser from "lucide-svelte/icons/circle-user";
  import { Key } from "lucide-svelte";

  const user = GetAuthContext(); 

  function seededLogin() {
    SetAuthContext({id: BigInt(0), name: "test_user", src: "https://cdn.britannica.com/08/190808-050-CB26C47B/The-Powerpuff-Girls-Bubbles-Blossom-Buttercup.jpg"});
  }

</script>

{#if $user !== undefined} 
    <div>
        <p>{$user.name}</p>
        <img src={$user.src} alt="user {$user.name}"/>
    </div>
{:else}
    <button onclick={seededLogin}>Login</button>
{/if}

    
  




