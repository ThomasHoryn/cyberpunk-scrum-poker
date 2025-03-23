<script lang="ts">
  import Footer from "./components/Footer.svelte";
  import { createRoom } from "./api/rooms.svelte";
  import type { CreateRoomRequest } from "./api/rooms.svelte";

  let roomName = $state("");
  let roomId = $state("");
  let error = $state("");
  let copySuccess = $state(false);

  let roomNameError = $state("");

  function validateRoomName() {
    if (roomName.trim().length < 3) {
      roomNameError = "DATA INSUFFICIENT: Name requires minimum 3 characters";
    } else if (roomName.trim().length > 30) {
      roomNameError = "DATA OVERFLOW: Exceeds permitted character limit";
    } else {
      roomNameError = "";
    }
  }

  async function handleCreateRoom(): Promise<void> {
    if (!roomName.trim()) {
      error =
        "ERR//0x734: INPUT VOID - Enter a valid room identifier, choomba!";
      // Add screen shake animation
      document
        .querySelector(".bg-light-gray")
        ?.classList.add("cyber-error-shake");
      setTimeout(() => {
        document
          .querySelector(".bg-light-gray")
          ?.classList.remove("cyber-error-shake");
      }, 500);
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
      // Add cyberpunk-themed prefix to error messages
      error = `SYS_FAILURE//0x${Math.floor(Math.random() * 999)
        .toString(16)
        .padStart(3, "0")}: ${response.error}`;
    } else if (response.data) {
      roomId = response.data.id;
      // Add success animation
      document
        .querySelector(".bg-light-gray")
        ?.classList.add("cyber-success-flash");
      setTimeout(() => {
        document
          .querySelector(".bg-light-gray")
          ?.classList.remove("cyber-success-flash");
      }, 1000);
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
                class="w-full bg-dark-gray text-neon-yellow border {roomName.trim()
                  ? roomName.trim().length < 3
                    ? 'border-neon-pink'
                    : 'border-neon-blue'
                  : 'border-neon-blue'} p-3 clip-corners-sm font-mono"
                maxlength="50"
                oninput={validateRoomName}
              />

              {#if roomNameError && roomName.trim()}
                <div class="mt-2 cyber-error-container">
                  <div class="cyber-error-glitch"></div>
                  <div class="cyber-error-content">
                    <div class="mr-2 cyber-error-icon">!</div>
                    <p class="digital-text text-sm error-message">
                      {roomNameError}
                    </p>
                  </div>
                </div>
              {/if}
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

        <!-- Error Message -->
        {#if error}
          <div
            class="mt-6 p-4 bg-dark-gray border-2 border-neon-pink clip-corners-sm text-neon-pink cyber-error-panel"
          >
            <div class="flex items-center">
              <div class="cyber-error-warning mr-2">!</div>
              <span class="font-bold digital-text">ERROR_DETECTED:</span>
            </div>
            <div class="mt-2 font-mono relative overflow-hidden">
              <span class="error-text">{error}</span>
              <div
                class="absolute top-0 left-0 w-full h-full cyber-error-scanline"
              ></div>
            </div>
            <div class="cyber-error-corner-tl"></div>
            <div class="cyber-error-corner-br"></div>
          </div>
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

  .clip-corners-lg {
    clip-path: polygon(
      0 20px,
      20px 0,
      calc(100% - 20px) 0,
      100% 20px,
      100% calc(100% - 20px),
      calc(100% - 20px) 100%,
      20px 100%,
      0 calc(100% - 20px)
    );
  }

  .cyber-button-sm {
    position: relative;
    overflow: hidden;
    clip-path: polygon(
      0 0,
      calc(100% - 5px) 0,
      100% 5px,
      100% 100%,
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

  .cyber-button-primary::after {
    content: "";
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    background: linear-gradient(
      to bottom,
      rgba(255, 255, 255, 0.1),
      transparent
    );
    opacity: 0;
    transition: opacity 0.2s ease;
  }

  .cyber-button-primary:active::after {
    opacity: 0.3;
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

  .cyber-error-container {
    position: relative;
    overflow: hidden;
    padding: 8px;
  }

  .cyber-error-container::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: linear-gradient(
      90deg,
      transparent,
      var(--color-neon-pink),
      transparent
    );
    opacity: 0.2;
    animation: error-pulse 2s ease-in-out infinite;
  }

  .cyber-error-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 18px;
    height: 18px;
    color: var(--color-black);
    background-color: var(--color-neon-pink);
    clip-path: polygon(50% 0%, 100% 38%, 82% 100%, 18% 100%, 0% 38%);
    font-weight: bold;
    font-size: 12px;
    animation: error-blink 1s ease-in-out infinite;
  }

  .cyber-error-container {
    position: relative;
    overflow: hidden;
    padding: 0.5rem;
    background: rgba(255, 0, 60, 0.1);
    border: 1px solid var(--color-neon-pink);
    animation: container-glitch 0.5s ease-out;
  }

  .cyber-error-glitch {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: linear-gradient(
      45deg,
      transparent 65%,
      var(--color-neon-pink) 65% 70%,
      transparent 70% 80%,
      var(--color-neon-blue) 80% 85%,
      transparent 85%
    );
    filter: brightness(1.2) contrast(2);
    mix-blend-mode: color-dodge;
    opacity: 0.7;
    animation: glitch-slide 3s linear infinite;
  }

  .cyber-error-content {
    position: relative;
    z-index: 1;
    display: flex;
    align-items: center;
  }

  .cyber-error-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 20px;
    height: 20px;
    background-color: var(--color-neon-pink);
    color: var(--color-black);
    clip-path: polygon(50% 0%, 100% 38%, 82% 100%, 18% 100%, 0% 38%);
    font-weight: bold;
    font-size: 14px;
    animation: icon-pulse 1s ease-in-out infinite alternate;
  }

  .error-message {
    margin-left: 0.5rem;
    color: var(--color-neon-pink);
    text-shadow: 0 0 5px var(--color-neon-pink);
    animation: text-glitch 3s infinite;
  }

  @keyframes container-glitch {
    0% {
      transform: translateX(-10px) skew(10deg);
      opacity: 0;
    }
    100% {
      transform: translateX(0) skew(0);
      opacity: 1;
    }
  }

  @keyframes glitch-slide {
    0% {
      transform: translateX(-100%) skew(-45deg);
    }
    100% {
      transform: translateX(100%) skew(-45deg);
    }
  }

  @keyframes icon-pulse {
    0% {
      transform: scale(1);
    }
    100% {
      transform: scale(1.1);
    }
  }

  @keyframes text-glitch {
    0% {
      transform: translateX(0);
    }
    20% {
      transform: translateX(-2px);
    }
    40% {
      transform: translateX(2px);
    }
    60% {
      transform: translateX(-1px);
    }
    80% {
      transform: translateX(1px);
    }
    100% {
      transform: translateX(0);
    }
  }

  @keyframes error-pulse {
    0% {
      opacity: 0.2;
    }
    50% {
      opacity: 0.4;
    }
    100% {
      opacity: 0.2;
    }
  }

  @keyframes error-blink {
    0%,
    100% {
      opacity: 1;
    }
    50% {
      opacity: 0.7;
    }
  }

  @keyframes error-noise-anim {
    0% {
      clip: rect(3px, 9999px, 42px, 0);
    }
    5% {
      clip: rect(34px, 9999px, 12px, 0);
    }
    10% {
      clip: rect(60px, 9999px, 26px, 0);
    }
    15% {
      clip: rect(13px, 9999px, 23px, 0);
    }
    20% {
      clip: rect(29px, 9999px, 73px, 0);
    }
    25% {
      clip: rect(74px, 9999px, 42px, 0);
    }
    30% {
      clip: rect(0, 0, 0, 0);
    }
    35% {
      clip: rect(10px, 9999px, 44px, 0);
    }
    40% {
      clip: rect(84px, 9999px, 36px, 0);
    }
    45% {
      clip: rect(0, 0, 0, 0);
    }
    50% {
      clip: rect(79px, 9999px, 5px, 0);
    }
    55% {
      clip: rect(27px, 9999px, 60px, 0);
    }
    60% {
      clip: rect(0, 0, 0, 0);
    }
    65% {
      clip: rect(39px, 9999px, 57px, 0);
    }
    70% {
      clip: rect(0, 0, 0, 0);
    }
    75% {
      clip: rect(94px, 9999px, 15px, 0);
    }
    80% {
      clip: rect(58px, 9999px, 93px, 0);
    }
    85% {
      clip: rect(0, 0, 0, 0);
    }
    90% {
      clip: rect(67px, 9999px, 46px, 0);
    }
    95% {
      clip: rect(34px, 9999px, 29px, 0);
    }
    100% {
      clip: rect(63px, 9999px, 37px, 0);
    }
  }

  .cyber-error-container {
    position: relative;
    overflow: hidden;
    animation: error-fade-in 0.3s ease-out;
  }

  .cyber-error-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 18px;
    height: 18px;
    color: var(--color-black);
    background-color: var(--color-neon-pink);
    clip-path: polygon(50% 0%, 100% 38%, 82% 100%, 18% 100%, 0% 38%);
    font-weight: bold;
    font-size: 12px;
    animation: error-blink 1s ease-in-out infinite;
  }

  .error-message {
    animation: glitch-text 0.4s linear;
  }

  @keyframes error-fade-in {
    from {
      opacity: 0;
      transform: translateY(-10px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @keyframes error-blink {
    0%,
    100% {
      opacity: 1;
    }
    50% {
      opacity: 0.7;
    }
  }

  @keyframes glitch-text {
    0% {
      transform: translate(0);
    }
    20% {
      transform: translate(-2px, 2px);
    }
    40% {
      transform: translate(-2px, -2px);
    }
    60% {
      transform: translate(2px, 2px);
    }
    80% {
      transform: translate(2px, -2px);
    }
    100% {
      transform: translate(0);
    }
  }

  @keyframes error-shake {
    10%,
    90% {
      transform: translate3d(-1px, 0, 0);
    }
    20%,
    80% {
      transform: translate3d(2px, 0, 0);
    }
    30%,
    50%,
    70% {
      transform: translate3d(-4px, 0, 0);
    }
    40%,
    60% {
      transform: translate3d(4px, 0, 0);
    }
  }

  @keyframes success-flash {
    0% {
      box-shadow: 0 0 0 rgba(var(--color-neon-blue), 0);
    }
    20% {
      box-shadow: 0 0 25px rgba(var(--color-neon-blue), 0.8);
    }
    100% {
      box-shadow: 0 0 5px rgba(var(--color-neon-blue), 0.1);
    }
  }

  .cyber-error-panel {
    position: relative;
  }

  .cyber-error-warning {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 22px;
    height: 22px;
    background-color: var(--color-neon-pink);
    color: var(--color-black);
    clip-path: polygon(50% 0%, 100% 38%, 82% 100%, 18% 100%, 0% 38%);
    font-weight: bold;
  }

  .cyber-error-corner-tl {
    position: absolute;
    top: 0;
    left: 0;
    width: 10px;
    height: 10px;
    border-top: 2px solid var(--color-neon-pink);
    border-left: 2px solid var(--color-neon-pink);
  }

  .cyber-error-corner-br {
    position: absolute;
    bottom: 0;
    right: 0;
    width: 10px;
    height: 10px;
    border-bottom: 2px solid var(--color-neon-pink);
    border-right: 2px solid var(--color-neon-pink);
  }

  .cyber-error-scanline {
    background: linear-gradient(
      to bottom,
      transparent 0%,
      rgba(var(--color-neon-pink), 0.2) 50%,
      transparent 100%
    );
    height: 10px;
    width: 100%;
    animation: scanline 2s linear infinite;
  }

  .error-text {
    display: inline-block;
    animation: error-text-glitch 3s infinite;
  }

  @keyframes scanline {
    0% {
      transform: translateY(-100%);
    }
    100% {
      transform: translateY(100%);
    }
  }

  @keyframes error-text-glitch {
    0%,
    100% {
      opacity: 1;
      transform: none;
    }
    7% {
      opacity: 1;
      transform: skew(-0.5deg, -0.9deg);
    }
    10% {
      opacity: 1;
      transform: none;
    }
    27% {
      opacity: 1;
      transform: none;
    }
    30% {
      opacity: 1;
      transform: skew(0.8deg, -0.1deg);
    }
    35% {
      opacity: 1;
      transform: none;
    }
    74% {
      opacity: 1;
      transform: none;
    }
    75% {
      opacity: 1;
      transform: skew(-1deg, 0.3deg);
    }
    76% {
      opacity: 1;
      transform: none;
    }
    93% {
      opacity: 1;
      transform: none;
    }
    94% {
      opacity: 1;
      transform: skew(0.3deg, 0.5deg);
    }
    95% {
      opacity: 1;
      transform: none;
    }
  }
</style>
