<script lang="ts">
  import {type User, createUser} from '@lib/types/user';
  import { authForm } from '@lib/auth/auth.svelte'

//careful with this, the snippet defined below actually makes use of its naming convention 
//to access the 'username' and 'password' values from the
//"Username" and "Password" text label

  let user: User = $state({username: "", password: ""})

  let feedback: HTMLParagraphElement | null = $state(null);

  async function signup() {
    let res = await createUser(user);

    if (!res.ok) {
      feedback.style.color = "red"
    }
    else {
      feedback.style.color = "green"
      authForm.setHidden();
    }
    feedback.innerText = await res.text();

  }

</script>

{#if authForm.isVisible("signup")}
{#snippet input_field(v: string, t: string)}
  <label class="flex flex-col items-start my-2">
    <p class="text-sm my-1">{v}</p>

    <input type={t} name={v.toLowerCase()} class="border-zinc-200 border-b-2 mx-5 rounded" bind:value={() => user[v.toLowerCase()], (value) => {user[v.toLowerCase()] = value}}/>
  </label>

{/snippet}


<div class="flex flex-col w-80 h-100 absolute m-auto inset-24 bg-zinc-800 box-border p-5 rounded-xl shadows-2xl shadow-2xl">
  <button class="w-[2em] self-end place-items-center " onclick={() => authForm.setHidden()}>x</button>
  <form onsubmit={(e) => { e.preventDefault(); signup()}}>

    <h1 class="text-xl mb-2">Sign up</h1>

    <div class="box-border px-7 pt-4">
      {@render input_field('Username', 'text')}
      {@render input_field('Password', 'password')}
    </div>

    <input class="w-full bg-blue-600 rounded-sm mt-10 my-5" type="submit" value="sign up"/>
    <p bind:this={feedback}></p>

    <p>Already have an account?</p>
    <p><button onclick={() => authForm.setVisible("login")} class="underline" >Log in</button> now!</p>

  </form>
</div>
{/if}
