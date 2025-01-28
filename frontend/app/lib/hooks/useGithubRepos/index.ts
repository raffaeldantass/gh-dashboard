import { useState } from 'react';
import { handleSearch } from '@/app/lib/handleSearch';
import type { Repository, SearchFilters, UseGithubReposReturn } from '@/app/types/github';

const PER_PAGE = 10;

export function useGithubRepos(): UseGithubReposReturn {
  const [repositories, setRepositories] = useState<Repository[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0);
  const [currentFilters, setCurrentFilters] = useState<Omit<SearchFilters, 'page' | 'per_page'>>({
    query: '',
    organization: undefined
  });

  const fetchRepositories = async (
    filters: Omit<SearchFilters, 'page' | 'per_page'>, 
    page: number
  ) => {
    setIsLoading(true);
    setError(null);

    const { data, error: searchError } = await handleSearch({
      ...filters,
      page,
      per_page: PER_PAGE
    });
    
    if (searchError) {
      setError(searchError);
      setRepositories([]);
      setTotalPages(0);
    }
    
    if (data) {
      setRepositories(data);
      setTotalPages(data.total_pages);
    }

    setIsLoading(false);
  };

  const handleSearchSubmit = async (
    filters: Omit<SearchFilters, 'page' | 'per_page'>
  ) => {
    setCurrentPage(1);
    setCurrentFilters(filters);
    await fetchRepositories(filters, 1);
  };

  const handlePageChange = async (page: number) => {
    setCurrentPage(page);
    await fetchRepositories(currentFilters, page);
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  return {
    repositories,
    isLoading,
    error,
    currentPage,
    totalPages,
    handleSearchSubmit,
    handlePageChange
  };
}