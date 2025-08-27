

<form class="flex flex-col w-80 h-100 absolute m-auto inset-24 bg-zinc-800 box-border p-10 rounded-xl shadows-2xl shadow-2xl">
  <h1 class="text-xl mb-2">Log in</h1>
  <label class="flex flex-col items-start my-2">
    <p class="text-sm my-1">Name </p>
    <input type="text" name="name" class="border-zinc-200 border-b-2 mx-5 rounded"/>
  </label>

  <label class="flex flex-col items-start my-2">
    <p class="text-sm my-1">Password </p>
    <input type="password" name="password" class="border-zinc-200 border-b-2 mx-5 rounded"/>
  </label>
 
  <input class="w-full bg-blue-600 rounded-sm mt-10 my-5" type="submit" value="log in"/>
  <p>Don't have an account?</p>
  <p><a class="underline" href="#">Sign up</a> for one!</p>
</form>
