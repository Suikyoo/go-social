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
    <div class="profile">
        <p>{$user.name}</p>
        <img src={$user.src} alt="user {$user.name}"/>
    </div>
{:else}
    <button onclick={seededLogin}>Login</button>
{/if}

<style>

    img {
        object-fit: cover;
        width: 2.3em;
        aspect-ratio: 1/1;
        border-radius: 100%;
        outline-offset: 3px;
        outline: 5px solid var(--color-blue-700); 
    }

    p {
        padding-right: 1em;
    }

    .profile {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
    }

</style>

  




