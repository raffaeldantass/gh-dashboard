'use client';

import { useEffect } from 'react';
import Link from 'next/link';
import { useSearchParams } from 'next/navigation';
import { useGithubRepos } from '@/app/lib/hooks/useGithubRepos';
import ResultList from '@/app/components/ResultList';
import { TokenManager } from '@/app/lib/utils';

// Components
const AuthenticationRequired = () => (
  <div className="flex flex-col justify-center items-center h-screen space-y-4">
    <h2 className="text-xl font-semibold text-gray-800">Authentication Required</h2>
    <Link 
      href="/"
      className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors"
    >
      Go to Login
    </Link>
  </div>
);

const PageHeader = ({ onLogout }: { onLogout: () => void }) => (
  <div className="flex justify-between items-center mb-8">
    <div>
      <h1 className="text-3xl font-bold text-gray-900">Repositories</h1>
      <p className="text-sm text-gray-500 mt-1">View and manage your GitHub repositories</p>
    </div>
    <button 
      onClick={onLogout}
      className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors"
    >
      Log Out
    </button>
  </div>
);

// Main Page Component
export default function RepositoriesPage() {
  const searchParams = useSearchParams();
  const { 
    repositories,
    isLoading,
    error,
    currentPage,
    totalPages,
    handlePageChange,
    logout,
    fetchRepositories
  } = useGithubRepos();

  // Handle token from URL
  useEffect(() => {
    const token = searchParams.get('token');
    if (token) {
      TokenManager.set(token);
      // Clean URL without page reload
      window.history.replaceState({}, '', '/repositories');
      // Fetch repositories with new token
      fetchRepositories(1);
    }
  }, [searchParams, fetchRepositories]);

  // Check authentication state
  const isAuthenticated = TokenManager.get() !== null;
  
  if (!isAuthenticated && !isLoading) {
    return <AuthenticationRequired />;
  }

  return (
    <main className="min-h-screen bg-gray-50">
      <div className="container mx-auto px-4 py-8">
        <PageHeader onLogout={logout} />

        <div className="bg-white rounded-lg shadow py-8 px-4">
          <ResultList
            repositories={repositories}
            isLoading={isLoading}
            error={error}
            currentPage={currentPage}
            totalPages={totalPages}
            onPageChange={handlePageChange}
          />
        </div>
      </div>
    </main>
  );
}