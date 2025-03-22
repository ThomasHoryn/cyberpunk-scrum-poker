import { config } from '../config/app.svelte';

export const state = {
    isLoading: false,
    error: null
};

// Type definitions
export interface ApiResponse<T> {
    data?: T;
    error?: string;
    status: number;
}

export interface RequestOptions {
    method?: 'GET' | 'POST' | 'PUT' | 'DELETE';
    headers?: Record<string, string>;
    body?: any;
}

async function generateSignature(data: any): Promise<{ signature: string, timestamp: string }> {
    if (!config.security.hmacSecret) {
        throw new Error("HMAC secret is not set");
    }

    const timestamp = new Date().toISOString();
    const encoder = new TextEncoder();

    const key = await crypto.subtle.importKey(
        "raw",
        encoder.encode(config.security.hmacSecret),
        { name: "HMAC", hash: "SHA-256" },
        false,
        ["sign"]
    );

    // Ensure data is stringified and combined with timestamp
    const signatureData = encoder.encode(JSON.stringify(data) + timestamp);
    const signatureArray = await crypto.subtle.sign("HMAC", key, signatureData);

    const signature = Array.from(new Uint8Array(signatureArray))
        .map(b => b.toString(16).padStart(2, "0"))
        .join("");

    return { signature, timestamp };
}

export async function apiRequest<T>(
    endpoint: string,
    options: RequestOptions = {}
): Promise<ApiResponse<T>> {
    console.log('Starting apiRequest for endpoint:', endpoint);
    console.log('Options:', options);

    state.isLoading = true;
    state.error = null;

    try {
        const { method = 'GET', headers = {}, body } = options;
        const url = `${config.api.baseUrl}/${endpoint}`;
        console.log('Constructed URL:', url);

        let signature = '';
        let timestamp = new Date().toISOString();

        if (body) {
            try {
                const signatureData = await generateSignature(body);
                signature = signatureData.signature;
                timestamp = signatureData.timestamp;
            } catch (signatureError) {
                console.error('Error generating signature:', signatureError);
                throw signatureError;
            }
        }

        console.log('Request headers:', {
            'Content-Type': 'application/json',
            'X-API-Key': config.security.apiKey,
            'X-Timestamp': timestamp,
            'X-Signature': signature,
            ...headers
        });
        console.log('Request body:', body ? JSON.stringify(body) : undefined);

        const response = await fetch(url, {
            method,
            headers: {
                'Content-Type': 'application/json',
                'X-API-Key': config.security.apiKey,
                'X-Timestamp': timestamp,
                'X-Signature': signature,
                ...headers
            },
            body: body ? JSON.stringify(body) : undefined
        });
        console.log('Response status:', response.status);

        const data = await response.json();
        console.log('Response data:', data);

        const result = {
            data: response.ok ? data : undefined,
            error: !response.ok ? data.error || 'Unknown error' : undefined,
            status: response.status
        };
        console.log('Returning result:', result);
        return result;
    } catch (err) {
        console.error('Error in apiRequest:', err);
        return {
            error: err instanceof Error ? err.message : 'Unknown error',
            status: 500
        };
    } finally {
        state.isLoading = false;
        console.log('apiRequest completed. isLoading set to false.');
    }
}

