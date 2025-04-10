<script lang="ts">
  import Footer from "./components/Footer.svelte";
  import { createRoom } from "./api/rooms.svelte";
  import type { CreateRoomRequest } from "./api/rooms.svelte";
  import { fade } from "svelte/transition";

  let roomName = $state("");
  let roomId = $state("");
  let error = $state("");
  let copySuccess = $state(false);

  async function handleCreateRoom(): Promise<void> {
    if (!roomName.trim()) {
      error =
        "ERR//0x734: INPUT VOID - Enter a valid room identifier, choomba!";
      return;
    }

    if (roomName.trim().length < 3) {
      error =
        "ERR//0x735: DATA INSUFFICIENT - Room ID requires minimum 3 characters to register in the net.";
      return;
    }

    error = "";
    roomId = "";
    copySuccess = false;

    const request: CreateRoomRequest = { name: roomName };
    const response = await createRoom(request);

    if (response.error) {
      error = `SYS_FAILURE//0x${Math.floor(Math.random() * 999)
        .toString(16)
        .padStart(3, "0")}: ${response.error}`;
    } else if (response.data) {
      roomId = response.data.id;
    }
  }

  function copyRoomId() {
    if (roomId) {
      navigator.clipboard.writeText(roomId);
      copySuccess = true;
      setTimeout(() => (copySuccess = false), 3000);
    }
  }
</script>

<div class="flex min-h-screen relative bg-dark-gray overflow-hidden">
  <!-- Cyberpunk Pattern Background -->
  <div
    class="absolute inset-0 pointer-events-none opacity-5 z-0 cyber-grid"
  ></div>

  <!-- Content Column -->
  <div
    class="w-full bg-dark-gray text-neon-yellow flex items-center justify-center p-6 md:p-8 relative z-10"
  >
    <div class="text-center space-y-8 max-w-md">
      <h1 class="text-3xl md:text-4xl font-sans">
        <span class="block glitch-text" data-text="Cyberpunk">Cyberpunk</span>
        <span class="block text-neon-pink">Scrum Poker</span>
      </h1>

      <div
        class="bg-light-gray p-6 border border-neon-blue shadow-neon clip-corners relative"
      >
        {#if !roomId}
          <p class="mb-6 digital-text">
            Create a room and invite your team to the voting session:
          </p>
          <!-- Error Message -->
          {#if error}
            <div
              class="mt-6 mb-8 p-4 bg-dark-gray border-2 border-neon-pink clip-corners-sm text-neon-pink cyber-error-panel"
              transition:fade={{ duration: 300 }}
            >
              <div class="flex items-center">
                <div class="cyber-error-warning mr-2">!</div>
                <span class="font-bold digital-text">ERROR_DETECTED:</span>
              </div>
              <div class="mt-2 font-mono relative overflow-hidden">
                <span class="error-text">{error}</span>
                <div
                  class="absolute top-0 left-0 w-full h-6 cyber-error-scanline"
                ></div>
              </div>
              <div class="cyber-error-corner-tl"></div>
              <div class="cyber-error-corner-br"></div>
            </div>
          {/if}

          <!-- Input for room name -->
          <div class="relative mb-6">
            <div
              class="absolute -inset-1 bg-gradient-to-r from-neon-pink via-neon-blue to-neon-pink opacity-50 blur-sm clip-corners-sm"
            ></div>
            <div class="relative">
              <input
                type="text"
                bind:value={roomName}
                placeholder="Enter room name"
                class="w-full bg-dark-gray text-neon-yellow border border-neon-blue p-3 clip-corners-sm font-mono"
                maxlength="50"
              />
            </div>
          </div>

          <!-- Create Room Button -->
          <button
            onclick={handleCreateRoom}
            class="cyber-button-primary bg-neon-pink text-black cursor-pointer font-bold py-3 px-6 text-lg hover:bg-neon-yellow disabled:opacity-50 disabled:cursor-not-allowed w-full"
          >
            Create Room
          </button>
        {:else}
          <!-- Room ID Display -->
          <div
            class="mt-6 p-4 bg-dark-gray border border-neon-yellow clip-corners-sm"
          >
            <p class="text-neon-yellow mb-2 text-xs">ROOM ID:</p>
            <div class="flex items-center justify-between">
              <span class="text-neon-pink font-mono tracking-wider"
                >{roomId}</span
              >
              <button
                onclick={copyRoomId}
                class="cyber-button-sm bg-neon-blue text-black px-3 py-1 text-xs hover:bg-neon-yellow"
              >
                Copy
              </button>
            </div>
          </div>
          {#if copySuccess}
            <p class="text-neon-green text-sm mt-2">
              Room ID copied to clipboard!
            </p>
          {/if}
        {/if}

        <!-- Tech decoration elements -->
        <div
          class="absolute top-2 right-2 w-6 h-6 border-t border-r border-neon-pink"
        ></div>
        <div
          class="absolute bottom-2 left-2 w-6 h-6 border-b border-l border-neon-blue"
        ></div>
      </div>

      <!-- Data streaming animation -->
      <div class="binary-stream">
        <div class="stream-line text-xs opacity-20">
          01001001010100110100111001
        </div>
        <div class="stream-line-2 text-xs opacity-15">
          10100111000110101100100111
        </div>
      </div>
    </div>
  </div>
</div>
<Footer />

<style>
  /* Cyberpunk styling */
  .clip-corners {
    clip-path: polygon(
      0 10px,
      10px 0,
      calc(100% - 10px) 0,
      100% 10px,
      100% calc(100% - 10px),
      calc(100% - 10px) 100%,
      10px 100%,
      0 calc(100% - 10px)
    );
  }

  .clip-corners-sm {
    clip-path: polygon(
      0 5px,
      5px 0,
      calc(100% - 5px) 0,
      100% 5px,
      100% calc(100% - 5px),
      calc(100% - 5px) 100%,
      5px 100%,
      0 calc(100% - 5px)
    );
  }

  .cyber-button-primary {
    position: relative;
    overflow: hidden;
    clip-path: polygon(
      0 10px,
      10px 0,
      calc(100% - 0px) 0,
      100% 0px,
      100% calc(100% - 0px),
      calc(100% - 10px) 100%,
      0px 100%,
      0 calc(100% - 10px)
    );
    transition: all 0.2s var(--ease-snappy);
  }

  .cyber-button-primary:hover {
    transform: translateY(-2px);
    box-shadow: 0 0 15px currentColor;
  }

  .cyber-button-primary:active {
    transform: translateY(3px) scale(0.97);
    box-shadow: 0 0 8px currentColor inset;
    background-image: linear-gradient(
      to bottom,
      rgba(0, 0, 0, 0.2),
      transparent 80%
    );
  }

  .shadow-neon {
    box-shadow:
      0 0 10px var(--color-neon-blue),
      0 0 20px var(--color-neon-pink);
  }

  .cyber-grid {
    background-size: 50px 50px;
    background-image: linear-gradient(
        to right,
        rgba(149, 217, 244, 0.1) 1px,
        transparent 1px
      ),
      linear-gradient(to bottom, rgba(149, 217, 244, 0.1) 1px, transparent 1px);
  }

  .glitch-text {
    position: relative;
    display: inline-block;
  }

  .glitch-text::before,
  .glitch-text::after {
    content: attr(data-text);
    position: absolute;
    width: 100%;
    height: 100%;
    left: 0;
    top: 0;
    opacity: 0.8;
  }

  .glitch-text::before {
    animation: glitch-anim-1 2s infinite linear alternate-reverse;
    clip: rect(44px, 450px, 56px, 0);
    left: 1px;
    text-shadow: -2px 0 var(--color-neon-pink);
  }

  .glitch-text::after {
    animation: glitch-anim-2 3.5s infinite linear alternate-reverse;
    left: -2px;
    clip: rect(44px, 450px, 56px, 0);
    text-shadow: 2px 0 var(--color-neon-blue);
  }

  .digital-text {
    position: relative;
    font-family: "Orbitron", sans-serif;
    letter-spacing: 0.5px;
  }

  /* Binary stream animation */
  .binary-stream {
    position: relative;
    height: 40px;
    overflow: hidden;
    font-family: monospace;
    letter-spacing: 2px;
  }

  .stream-line {
    position: absolute;
    white-space: nowrap;
    animation: stream-animation 20s linear infinite;
    color: var(--color-neon-yellow);
  }

  .stream-line-2 {
    position: absolute;
    white-space: nowrap;
    animation: stream-animation-reverse 15s linear infinite;
    color: var(--color-neon-pink);
    top: 20px;
  }

  @keyframes stream-animation {
    0% {
      transform: translateX(-100%);
    }
    100% {
      transform: translateX(100%);
    }
  }

  @keyframes stream-animation-reverse {
    0% {
      transform: translateX(100%);
    }
    100% {
      transform: translateX(-100%);
    }
  }

  @keyframes glitch-anim-1 {
    0% {
      clip: rect(42px, 9999px, 26px, 0);
    }
    10% {
      clip: rect(33px, 9999px, 19px, 0);
    }
    20% {
      clip: rect(38px, 9999px, 65px, 0);
    }
    30% {
      clip: rect(0, 0, 0, 0);
    }
    40% {
      clip: rect(0, 0, 0, 0);
    }
    50% {
      clip: rect(62px, 9999px, 78px, 0);
    }
    60% {
      clip: rect(0, 0, 0, 0);
    }
    70% {
      clip: rect(14px, 9999px, 26px, 0);
    }
    80% {
      clip: rect(0, 0, 0, 0);
    }
    90% {
      clip: rect(56px, 9999px, 97px, 0);
    }
    100% {
      clip: rect(24px, 9999px, 40px, 0);
    }
  }

  @keyframes glitch-anim-2 {
    0% {
      clip: rect(36px, 9999px, 100px, 0);
    }
    15% {
      clip: rect(0, 0, 0, 0);
    }
    20% {
      clip: rect(40px, 9999px, 91px, 0);
    }
    35% {
      clip: rect(0, 0, 0, 0);
    }
    40% {
      clip: rect(19px, 9999px, 7px, 0);
    }
    55% {
      clip: rect(0, 0, 0, 0);
    }
    60% {
      clip: rect(55px, 9999px, 63px, 0);
    }
    75% {
      clip: rect(0, 0, 0, 0);
    }
    80% {
      clip: rect(14px, 9999px, 39px, 0);
    }
    95% {
      clip: rect(0, 0, 0, 0);
    }
    100% {
      clip: rect(82px, 9999px, 17px, 0);
    }
  }
</style>
