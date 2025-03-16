<script lang="ts">
  import HomePage from './lib/HomePage.svelte';
  import TestTailwind from './lib/TestTailwind.svelte';
  
  // Set to true to test Tailwind
  const showTest = false;
</script>

<main>
  {#if showTest}
    <TestTailwind />
  {:else}
    <HomePage />
  {/if}
</main>
