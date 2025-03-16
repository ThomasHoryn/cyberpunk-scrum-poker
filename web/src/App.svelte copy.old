<script lang="ts">
  let roomName: string = "";
  let result: string = "";
  let isLoading: boolean = false;

  // Configuration
  const API_KEY: string = "cyberpunk_dev_2025";
  const HMAC_SECRET: string = "neonlights2077";
  const API_URL: string = "http://localhost:8054/api/v1/rooms";

  interface ApiResponse {
    id?: string;
    error?: string;
  }

  async function createRoom(): Promise<void> {
    isLoading = true;
    result = "";
    console.log('Raw Input:', JSON.stringify({ name: roomName }));
    
    try {
      const timestamp: string = new Date().toISOString();
      const requestBody = { name: roomName };
      
      // Generate HMAC signature
      const encoder = new TextEncoder();
      const key = await crypto.subtle.importKey(
        "raw",
        encoder.encode(HMAC_SECRET),
        { name: "HMAC", hash: "SHA-256" },
        false,
        ["sign"]
      );
      
      const data = encoder.encode(JSON.stringify(requestBody) + timestamp);
      const signatureArray = await crypto.subtle.sign("HMAC", key, data);
      const signature = Array.from(new Uint8Array(signatureArray))
        .map(b => b.toString(16).padStart(2, "0"))
        .join("");

console.log("Generated Signature:", signature);
console.log("Timestamp:", timestamp);
console.log("Request Body:", requestBody);
console.log("API Key:", API_KEY);


      // Send request
      const response: Response = await fetch(API_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "X-API-Key": API_KEY,
          "X-Timestamp": timestamp,
          "X-Signature": signature
        },
        body: JSON.stringify(requestBody)
      });

      const responseData: ApiResponse = await response.json();
      result = response.ok 
        ? `Room created! ID: ${responseData.id}`
        : `Error: ${responseData.error || "Unknown error"}`;

    } catch (err: unknown) {
      result = `Error: ${err instanceof Error ? err.message : 'Unknown error'}`;
    } finally {
      isLoading = false;
    }
  }
</script>

<main>
  <h1>Test Room Creation</h1>
  
  <input 
    type="text" 
    bind:value={roomName}
    placeholder="Enter room name"
    minlength="3"
    maxlength="50"
  >
  
  <button on:click={createRoom} disabled={isLoading}>
    {isLoading ? 'Creating...' : 'Create Room'}
  </button>

  {#if result}
    <div class:success={result.includes('created')} class:error={result.includes('Error')}>
      {result}
    </div>
  {/if}
</main>

<style>
  main {
    padding: 2rem;
    font-family: system-ui;
    max-width: 600px;
    margin: 0 auto;
  }
  
  input, button {
    padding: 0.5rem;
    margin-right: 1rem;
    margin-bottom: 1rem;
  }
  
  .success { color: green; }
  .error { color: red; }
</style>
