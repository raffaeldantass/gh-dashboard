import { useState, useEffect } from 'react';
import { fetchApi } from '@/app/lib/api';
import { TokenManager } from '@/app/lib/utils';
import type { 
  Repository, 
  UseGithubReposReturn, 
  PaginatedResponse 
} from '@/app/types/github';
import { useRouter } from 'next/navigation';

export function useGithubRepos(): UseGithubReposReturn {
  const router = useRouter();
  
  // State management
  const [repositories, setRepositories] = useState<Repository[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);

  // Initial load
  useEffect(() => {
    const token = TokenManager.get();
    if (token) {
      TokenManager.set(token);
      fetchRepositories(1);
    }
  }, []);

  // API interaction
  const fetchRepositories = async (page: number): Promise<void> => {
    setIsLoading(true);
    setError(null);
    
    try {
      const token = TokenManager.get();
      if (!token) {
        throw new Error('No authentication token found');
      }

      const response = await fetchApi(`/get-repositories?page=${page}&per_page=10`, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });
      
      const data = response as PaginatedResponse;
      setRepositories(data?.repositories ?? []);
      setCurrentPage(data.current_page);
      setTotalPages(data.total_pages);
    } catch (err) {
      setError('Failed to fetch repositories');
      console.error('Error fetching repositories:', err);
    } finally {
      setIsLoading(false);
    }
  };

  // Event handlers
  const handlePageChange = async (page: number): Promise<void> => {
    await fetchRepositories(page);
  };

  const logout = (): void => {
    TokenManager.clear();
    setRepositories([]);
    setIsLoading(false);
    setError(null);
    setCurrentPage(1);
    setTotalPages(1);
    router.push('/');
  };

  return {
    repositories,
    isLoading,
    error,
    currentPage,
    totalPages,
    fetchRepositories,
    handlePageChange,
    logout,
  };
}
