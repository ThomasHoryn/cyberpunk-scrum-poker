<script lang="ts">
    import Love from "../icons/Love.svelte";

    let currentYear: number = new Date().getFullYear();
</script>

<div class="fixed bottom-0 left-0 right-0 z-10 pointer-events-none">
    <!-- Ultra thin separator line -->
    <div class="h-px w-full bg-gradient-to-r from-transparent via-neon-blue to-transparent opacity-20"></div>

    <!-- Main footer content - fully transparent -->
    <div class="relative py-2 px-4">
      <!-- Footer content with transparency -->
      <div class="flex flex-wrap items-center justify-between">
        <!-- Left content -->
        <div class="pointer-events-auto flex items-center my-1 opacity-60 hover:opacity-100 transition-opacity duration-300">
          <span class="text-neon-yellow text-xs">Made with &nbsp;</span>
          <div class="heart-icon-container">
            <Love size={18} />
          </div>
          <span class="text-neon-yellow text-xs">&nbsp; in Warsaw by</span>
          <a href="https://www.linkedin.com/in/tomaszhoryn"
             target="_blank"
             class="text-neon-pink ml-1 hover:text-neon-blue transition-colors duration-300 text-xs relative">
            Tomasz Horyn
          </a>
        </div>

        <!-- Right content - version/date -->
        <div class="pointer-events-auto opacity-40 hover:opacity-80 transition-opacity duration-300 my-1">
          <span class="text-neon-blue text-xs">
            CYBER-SCRUM-POKER v1.0 // {currentYear}
          </span>
        </div>
      </div>
    </div>
</div>

