<script lang="ts">
  import {type User, authenticateUser} from '@lib/types/user';
  import { authForm } from '@lib/auth/auth.svelte'
  import { userProfile } from '@components/Profile.svelte';

  import Input from '@components/Input.svelte';

//careful with this, the snippet defined below actually makes use of its naming convention 
//to access the 'username' and 'password' values from the
//"Username" and "Password" text label

  let user: User = $state({username: "", password: ""})

  let feedback: HTMLParagraphElement | null = $state(null);

  async function login() {
    let res = await authenticateUser(user);

    if (!res.ok) {
      feedback.innerText = await res.text();
      feedback.style.color = "red"
    }
    else {
      feedback.innerText = "succesfully signed in";
      feedback.style.color = "green"
      userProfile.username = user.username;
      authForm.setHidden();
    }

  }

</script>

{#if authForm.isVisible("login")}

<div class="flex flex-col w-80 h-100 absolute m-auto inset-24 bg-zinc-800 box-border p-5 rounded-xl shadows-2xl shadow-2xl">
  <button class="w-[2em] self-end place-items-center " onclick={() => authForm.setHidden()}>x</button>
  <form class="bg-inherit" onsubmit={(e) => { e.preventDefault(); login()}}>

    <h1 class="text-xl mb-2">Log in</h1>

    <div class="box-border px-7 pt-4 bg-inherit">

      <Input name="Username" type="text" bind:value={user.username}/>
      <Input name="Password" type="password" bind:value={user.password}/>
        <!--
      {@render input_field('Username', 'text')}
      {@render input_field('Password', 'password')}
        -->
    </div>

    <button class="w-full bg-blue-600 rounded-sm mt-10 my-5"> login </button>
    <p bind:this={feedback}></p>

    <p>Don't have an account?</p>
    <p><button onclick={() => authForm.setVisible("signup")} class="underline" >Sign up</button> for one!</p>

  </form>
</div>
{/if}
