<script lang="ts">
  import { authForm } from "@lib/auth/auth.svelte"
  import { type User } from "@lib/types/user"

  

</script>

<script module>
  interface UserProfile extends User {
    setUsername(value: string): void;

  }
  //do not set ID as the cookie already contains the user_id
  const userProfile: UserProfile = $state({
    username: localStorage.getItem("username") || "",
    setUsername: (value) => {
      localStorage.setItem("username", value);
      userProfile.username = value;
    }
  });
  export { userProfile };

</script>

{#if userProfile.username.length != 0}
    <div class="flex flex-row justify-between">
      <button onclick={() => {userProfile.setUsername(""); userProfile.setID("")}}>Log out</button>
      |
      <p>{userProfile.username}</p>
        <!-- add image -->
    </div>
{:else}
  <button onclick={() => authForm.setVisible("login")}>Login</button>
{/if}

    
  




