'use client';

import { useState } from 'react';
import SearchForm from '@/app/components/SearchForm';
import ResultList from '@/app/components/ResultList';
import { handleSearch } from '@/app/lib/handleSearch';
import type { Repository, SearchFilters } from '@/app/types/github';

export default function Home() {
  const [repositories, setRepositories] = useState<Repository[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const onSearch = async (filters: SearchFilters) => {
    setIsLoading(true);
    setError(null);

    const { data, error: searchError } = await handleSearch(filters);
    
    if (searchError) {
      setError(searchError);
    }
    
    if (data) {
      setRepositories(data);
    }

    setIsLoading(false);
  };

  return (
    <main className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-8">GitHub Repositories</h1>

      <SearchForm onSearch={onSearch} isLoading={isLoading} />
      <ResultList repositories={repositories} isLoading={isLoading} error={error} />
    </main>
  );
}