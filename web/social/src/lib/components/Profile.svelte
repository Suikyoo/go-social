<script lang="ts">
  import { authForm } from "@lib/auth/auth.svelte"
  import { type User } from "@lib/types/user"

  

</script>

<script module>
  interface UserProfile extends User {

    clear(): void;

  }
  const userProfile: UserProfile = $state({
    username: "",
    clear() {
      this.username = "";
    }
  });
  export { userProfile };

</script>

{#if userProfile.username.length != 0}
    <div class="flex flex-row">
      <button onclick={() => userProfile.clear()}>Log out</button>
      <p>{userProfile.username}</p>
        <!-- add image -->
    </div>
{:else}
  <button onclick={() => authForm.setVisible("login")}>Login</button>
{/if}

    
  




