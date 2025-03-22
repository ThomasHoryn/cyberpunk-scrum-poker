export const config = {
  api: {
    baseUrl: import.meta.env.VITE_API_URL || 'http://localhost:8054/api/v1',
    version: 'v1',
    timeout: 30000,
  },
  security: {
    apiKey: import.meta.env.VITE_API_KEY,
    hmacSecret: import.meta.env.VITE_HMAC_SECRET,
  },
  app: {
    name: 'Cyberpunk Scrum Poker',
    version: '1.0.0',
  }
};
