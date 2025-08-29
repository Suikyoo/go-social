<script lang="ts">
  import {type User, authenticateUser} from '../types/user';

//careful with this, the snippet defined below actually makes use of its naming convention 
//to access the 'username' and 'password' values from the
//"Username" and "Password" text label

  let user: User = $state({username: "", password: ""})

  let feedback: {status: boolean, content: string} = $state({status: false, content: ""});

  async function login() {
    let res = await authenticateUser(user);

    if (!res.ok) {
      feedback.status = false;
      feedback.content = await res.text();
    }
    else {
      feedback.status = true;
      feedback.content = "successfully signed in";
    }

  }

</script>

{#snippet input_field(v: string, t: string)}
  <label class="flex flex-col items-start my-2">
    <p class="text-sm my-1">{v}</p>

    <input type={t} name={v.toLowerCase()} class="border-zinc-200 border-b-2 mx-5 rounded" bind:value={() => user[v.toLowerCase()], (value) => {user[v.toLowerCase()] = value}}/>
  </label>

{/snippet}

<form  class="flex flex-col w-80 h-100 absolute m-auto inset-24 bg-zinc-800 box-border p-10 rounded-xl shadows-2xl shadow-2xl">

  <h1 class="text-xl mb-2">Log in</h1>

  {@render input_field('Username', 'text')}
  {@render input_field('Password', 'password')}
  <input class="w-full bg-blue-600 rounded-sm mt-10 my-5" type="submit" value="log in"/>

  <p>Don't have an account?</p>
  <p><a onclick={() => console.log($state.snapshot(user))} class="underline" href="#">Sign up</a> for one!</p>

</form>
