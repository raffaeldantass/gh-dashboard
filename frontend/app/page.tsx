'use client';

import React from 'react';
import SearchForm from '@/app/components/SearchForm';
import ResultList from '@/app/components/ResultList';
import Pagination from '@/app/components/Pagination';
import { useGithubRepos } from '@/app/lib/hooks/useGithubRepos';

export default function Home() {
  const {
    repositories,
    isLoading,
    error,
    currentPage,
    totalPages,
    handleSearchSubmit,
    handlePageChange
  } = useGithubRepos();

  return (
    <main className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-8">GitHub Repositories</h1>
      <SearchForm onSearch={handleSearchSubmit} isLoading={isLoading} />
      <ResultList repositories={repositories} isLoading={isLoading} error={error} />
      <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={handlePageChange}
        isLoading={isLoading}
      />
    </main>
  );
}