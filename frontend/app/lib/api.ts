// frontend/app/lib/api.ts
const BASE_URL = '/api';

export async function fetchApi(endpoint: string, options: RequestInit = {}) {
  const response = await fetch(`${BASE_URL}${endpoint}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
    credentials: 'include', // Important for cookies
  });

  if (!response.ok) {
    const error = await response.json().catch(() => ({}));
    throw new Error(error.message || `API call failed: ${response.statusText}`);
  }

  return response.json();
}