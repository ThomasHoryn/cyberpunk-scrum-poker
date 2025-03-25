<script lang="ts">
  import { currentPage } from "./lib/state.svelte";
  import HomePage from "./lib/HomePage.svelte";
  import NotFound from "./lib/NotFound.svelte";
  import Showcase from "./lib/designSystem/Showcase.svelte";
</script>

<main>
  {#if $currentPage === "home"}
    <HomePage />
  {:else if $currentPage === "showcase"}
    <Showcase />
  {:else}
    <NotFound />
  {/if}
</main>
