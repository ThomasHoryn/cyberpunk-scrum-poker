<script lang="ts">
  let activeTab = $state('backgrounds');
  
  const sections = [
    { id: 'backgrounds', label: 'Background Styles' },
    { id: 'navigation', label: 'Navigation Systems' }
  ];
</script>

<div class="min-h-screen bg-dark-gray relative overflow-hidden">
  <div class="absolute inset-0 cyber-grid animate-grid-pulse opacity-10"></div>
  <div class="fixed inset-0 scanline-overlay pointer-events-none"></div>

  <div class="relative z-10 container mx-auto px-4 py-16">
    <nav class="mb-16">
      <ul class="flex space-x-8">
        {#each sections as section}
          <li>
            <button
              class="cyber-nav-item {activeTab === section.id ? 'active' : ''}"
              onclick={() => activeTab = section.id}
            >
              {section.label}
            </button>
          </li>
        {/each}
      </ul>
    </nav>

    {#if activeTab === 'backgrounds'}
      <section class="space-y-12">
        <h2 class="text-3xl text-neon-pink mb-8">Background Patterns</h2>
        
        <div class="cyber-pattern-card">
          <div class="cyber-grid-demo animate-grid-move"></div>
          <div class="pattern-info">
            <h3 class="text-2xl text-neon-blue">Cyber Grid</h3>
            <code class="text-neon-yellow">class="cyber-grid"</code>
          </div>
        </div>

        <div class="cyber-pattern-card">
          <div class="scanline-demo"></div>
          <div class="pattern-info">
            <h3 class="text-2xl text-neon-blue">Scanline Overlay</h3>
            <code class="text-neon-yellow">class="scanline-overlay"</code>
          </div>
        </div>
      </section>
    {/if}

    {#if activeTab === 'navigation'}
      <section class="space-y-12">
        <h2 class="text-3xl text-neon-blue mb-8">Navigation Components</h2>
        
        <div class="cyber-nav-demo">
          <nav class="cyber-nav-bar">
            <a href="#dashboard" class="nav-item active">Dashboard</a>
            <a href="#components" class="nav-item">Components</a>
            <a href="#styleguide" class="nav-item">Styleguide</a>
          </nav>
        </div>
      </section>
    {/if}
  </div>
</div>

<style>
  .cyber-grid {
    background-image: linear-gradient(to right, var(--color-neon-blue) 1px, transparent 1px),
                      linear-gradient(to bottom, var(--color-neon-pink) 1px, transparent 1px);
    background-size: 50px 50px;
  }

  .animate-grid-pulse {
    animation: grid-pulse 20s infinite linear;
  }

  .scanline-overlay {
    background: repeating-linear-gradient(
      0deg,
      transparent 0px,
      transparent 1px,
      rgba(0, 255, 255, 0.05) 2px,
      transparent 3px
    );
  }

  .cyber-nav-item {
    position: relative;
    padding: 0.75rem 1.5rem;
    border: 2px solid var(--color-neon-blue);
    color: var(--color-neon-yellow);
    clip-path: polygon(10px 0, 100% 0, calc(100% - 10px) 100%, 0 100%);
    transition: all 0.3s ease;
  }
  
  .cyber-nav-item:hover {
    border-color: var(--color-neon-pink);
    color: var(--color-neon-pink);
    box-shadow: 0 0 15px var(--color-neon-pink);
  }
  
  .cyber-nav-item.active {
    background-color: var(--color-neon-pink);
    color: black;
    border-color: var(--color-neon-pink);
  }

  @keyframes grid-pulse {
    0% { opacity: 0.1; }
    50% { opacity: 0.15; }
    100% { opacity: 0.1; }
  }

  @keyframes grid-move {
    from { background-position: 0 0; }
    to { background-position: 100px 100px; }
  }

  .cyber-pattern-card {
    background-color: var(--color-dark-gray);
    border: 2px solid var(--color-neon-blue);
    padding: 1rem;
    margin-bottom: 1rem;
  }

  .cyber-grid-demo, .scanline-demo {
    height: 200px;
    width: 100%;
    margin-bottom: 1rem;
  }

  .cyber-grid-demo {
    background-image: inherit;
    background-size: inherit;
  }

  .scanline-demo {
    background: inherit;
  }

  .pattern-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .cyber-nav-bar {
    display: flex;
    gap: 1rem;
    padding: 0.5rem;
    background-color: var(--color-dark-gray);
    border: 2px solid var(--color-neon-blue);
  }

  .nav-item {
    padding: 0.5rem 1rem;
    color: var(--color-neon-yellow);
    transition: all 0.3s ease;
  }

  .nav-item:hover, .nav-item.active {
    background-color: var(--color-neon-blue);
    color: black;
  }
</style>
