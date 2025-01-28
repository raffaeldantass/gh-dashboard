import type { Repository, SearchFilters } from '@/app/types/github';

interface SearchResponse {
  data: Repository[] | null;
  error: string | null;
  total_pages?: number;
}

export async function handleSearch(filters: SearchFilters): Promise<SearchResponse> {
  try {
    const response = await fetch('/api/repositories', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(filters),
    });

    if (!response.ok) {
      throw new Error(`Request error - status: ${response.status}`); // This could be logged to a service like Sentry
    }

    const data = await response.json();
    return { data, error: null };
  } catch (error) {
    console.error('Error fetching repositories:', error); // This could be logged to a service like Sentry
    return { 
      data: null, 
      error: error instanceof Error ? error.message : 'An unknown error occurred' 
    };
  }
}