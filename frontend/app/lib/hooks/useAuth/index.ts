import { useEffect, useState } from 'react';
import { useSearchParams } from 'next/navigation';
import { fetchApi } from '@/app/lib/api';

interface Repository {
  id: number;
  name: string;
}

export function useAuth() {
  const [token, setToken] = useState<string | null>(null);
  const [repositories, setRepositories] = useState<Repository[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const searchParams = useSearchParams();

  console.log('repositories', repositories);


  useEffect(() => {
    const urlToken = searchParams.get('token');
    const storedToken = sessionStorage.getItem('github_token');

      if (urlToken) {
        sessionStorage.setItem('github_token', urlToken);
        setToken(urlToken);
        return;
      }
      
      if (storedToken) {
        setToken(storedToken);
      }
  }, [searchParams]);

  const fetchRepositories = async () => {
    if (!token) return;

    setLoading(true);
    setError(null);
    
    try {
      const data = await fetchApi('/repositories', {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });
      setRepositories(data);
    } catch (err) {
      setError('Failed to fetch repositories');
      console.error('Error fetching repositories:', err);
    } finally {
      setLoading(false);
    }
  };

  const logout = () => {
    sessionStorage.removeItem('github_token');
    setToken(null);
    setRepositories([]);
  };

  return {
    token,
    isAuthenticated: !!token,
    repositories,
    loading,
    error,
    fetchRepositories,
    logout
  };
}