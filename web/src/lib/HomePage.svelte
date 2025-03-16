<script lang="ts">
  import { onMount } from 'svelte';
  
  let glitchText = '';
  const fullText = 'CYBERPUNK SCRUM POKER';
  let currentIndex = 0;
  let showTerminal = false;
  
  onMount(() => {
    // Simulate terminal loading
    setTimeout(() => {
      showTerminal = true;
      
      // Start typing animation
      const typeInterval = setInterval(() => {
        if (currentIndex < fullText.length) {
          glitchText = fullText.substring(0, currentIndex + 1);
          currentIndex++;
        } else {
          clearInterval(typeInterval);
        }
      }, 150);
      
      return () => clearInterval(typeInterval);
    }, 800);
  });
</script>

<div class="min-h-screen bg-cyber-black flex flex-col items-center justify-center p-4 relative overflow-hidden">
  <!-- Scanlines effect -->
  <div class="absolute inset-0 pointer-events-none after:content-[''] after:absolute after:top-0 after:left-0 after:right-0 after:bottom-0 after:bg-scanline"></div>
  
  <!-- Noise overlay -->
  <div class="absolute inset-0 bg-[url('/noise.jpeg')] opacity-5 mix-blend-overlay pointer-events-none"></div>
  
  <!-- CRT flicker effect -->
  <div class="absolute inset-0 bg-cyber-blue/5 animate-flicker pointer-events-none"></div>
  
  <!-- Main content -->
  <div class="relative z-10 w-full max-w-4xl mx-auto">
    {#if showTerminal}
      <!-- Terminal interface container -->
      <div class="border border-cyber-blue/70 bg-cyber-black/90 p-6 rounded relative">
        <!-- Terminal window controls -->
        <div class="flex gap-2 mb-4">
          <div class="w-3 h-3 rounded-full bg-cyber-red"></div>
          <div class="w-3 h-3 rounded-full bg-cyber-yellow"></div>
          <div class="w-3 h-3 rounded-full bg-cyber-green"></div>
        </div>
        
        <!-- Header with glitch effect -->
        <header class="text-center mb-16">
          <h1 class="text-5xl md:text-7xl font-cyberpunk font-bold mb-6 text-transparent bg-clip-text bg-gradient-to-r from-cyber-blue via-cyber-purple to-cyber-pink">
            {glitchText}<span class="animate-pulse text-cyber-blue">_</span>
          </h1>
          <p class="text-cyber-green font-glitch text-xl md:text-2xl">
            <span class="relative inline-block animate-glitch">PLANNING POKER FOR THE DIGITAL WASTELAND</span>
          </p>
          <div class="h-0.5 w-1/2 mx-auto mt-6 bg-gradient-to-r from-transparent via-cyber-blue to-transparent"></div>
        </header>
        
        <!-- Main actions -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8 mb-16">
          <!-- Create Room -->
          <div class="border border-cyber-blue bg-cyber-black/50 p-8 rounded hover:shadow-neon-blue transition-all duration-300 group">
            <h2 class="text-cyber-blue text-3xl font-cyberpunk mb-4 group-hover:animate-glow">CREATE ROOM</h2>
            <p class="text-gray-400 mb-6">Launch your own digital haven for estimation sessions.</p>
            <button class="w-full bg-gradient-to-r from-cyber-blue to-cyber-purple text-white font-bold py-3 px-6 rounded hover:from-cyber-purple hover:to-cyber-blue transition-all duration-300">
              NETRUN >>
            </button>
          </div>
          
          <!-- Join Room -->
          <div class="border border-cyber-pink bg-cyber-black/50 p-8 rounded hover:shadow-neon-pink transition-all duration-300 group">
            <h2 class="text-cyber-pink text-3xl font-cyberpunk mb-4 group-hover:animate-glow">JOIN ROOM</h2>
            <p class="text-gray-400 mb-6">Enter an existing neural network to collaborate.</p>
            <button class="w-full bg-gradient-to-r from-cyber-pink to-cyber-purple text-white font-bold py-3 px-6 rounded hover:from-cyber-purple hover:to-cyber-pink transition-all duration-300">
              CONNECT >>
            </button>
          </div>
        </div>
        
        <!-- Features -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-16">
          <div class="border-t-2 border-cyber-yellow pt-4">
            <h3 class="text-cyber-yellow text-xl font-cyberpunk mb-2">SECURE VOTING</h3>
            <p class="text-gray-400">End-to-end encrypted estimation, locked down like a corporate server.</p>
          </div>
          <div class="border-t-2 border-cyber-green pt-4">
            <h3 class="text-cyber-green text-xl font-cyberpunk mb-2">REAL-TIME SYNC</h3>
            <p class="text-gray-400">Instant neural connection with your team through netrunning protocols.</p>
          </div>
          <div class="border-t-2 border-cyber-red pt-4">
            <h3 class="text-cyber-red text-xl font-cyberpunk mb-2">CUSTOM DECKS</h3>
            <p class="text-gray-400">Hack the system with your own card designs and voting scales.</p>
          </div>
        </div>
        
        <!-- Footer -->
        <footer class="text-center text-gray-500 font-glitch text-sm">
          <div class="inline-block border border-gray-800 px-4 py-1 rounded-sm">
            <p>SYSTEM V2.0.7.7 :: SECURITY LEVEL: ALPHA :: CORP-ID: CSP-2025</p>
          </div>
        </footer>
      </div>
    {/if}
  </div>
</div>

<style>
  .bg-scanline {
    background: repeating-linear-gradient(
      to bottom,
      transparent 0px,
      transparent var(--scanline-gap),
      rgba(0, 240, 255, 0.1) var(--scanline-gap),
      rgba(0, 240, 255, 0.1) calc(var(--scanline-gap) + var(--scanline-height))
    );
    animation: scan 8s linear infinite;
  }
  
  @keyframes scan {
    0% { background-position: 0 0; }
    100% { background-position: 0 100vh; }
  }
</style>
