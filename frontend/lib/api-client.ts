import { useAuthStore } from '@/store/useAuthStore';

const BASE_URL = process.env.NEXT_PUBLIC_API_URL || '';

interface FetchOptions extends RequestInit {
  headers?: Record<string, string>;
}

export async function fetchClient<T>(url: string, options: FetchOptions = {}): Promise<T> {
  const token = useAuthStore.getState().token;
  
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    ...options.headers,
  };

  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  const response = await fetch(`${BASE_URL}${url}`, {
    ...options,
    headers,
  });

  if (!response.ok) {
    throw new Error(`HTTP Error: ${response.status}`);
  }

  const text = await response.text();
  let data;
  try {
    data = text ? JSON.parse(text) : {};
  } catch (error) {
    throw new Error('Invalid JSON response');
  }

  // Handle custom code from API if needed (e.g. code != 0)
  if (data.code !== undefined && data.code !== 0) {
      const errorMessage = data.message && data.message.trim() !== '' 
        ? data.message 
        : `Request failed with code ${data.code}`;
      throw new Error(errorMessage);
  }

  return data;
}
