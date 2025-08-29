<script lang="ts">
  import {House} from "@lucide/svelte";
  import {type Icon as IconType} from "@lucide/svelte";


  interface MenuItem {
    name: string;
    href: string;
    icon: typeof IconType;
  }

  const menuItems: MenuItem[] = [
    {
      name: "Home",
      href: "/",
      icon: House,
    }
  ]

</script>


<nav class="flex flex-row w-full lg:w-20 justify-evenly lg:justify-start lg:h-full lg:flex-col">

  {#each menuItems as item}
    {@const Icon = item.icon}
    <div class="p-5">
      <Icon strokeWidth={1.75} />
    </div>
  {/each}

</nav>
