import React from 'react';
import type { SearchFilters } from '@/app/types/github';

interface SearchFormProps {
  onSearch: (filters: SearchFilters) => void;
  isLoading: boolean;
}

const SearchForm: React.FC<SearchFormProps> = ({ onSearch, isLoading }) => {
  const [query, setQuery] = React.useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSearch({ query });
  };

  return (
    <div className="pb-10 mb-6 border-b border-gray-200">
      <div className="pt-6">
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="flex flex-col md:flex-row gap-4 md:items-center justify-center">
            <p className="h-auto m-0 p-0"> Type a valid Org or Personal account: </p>
            <div className="relative flex-1">
              <input
                type="text"
                placeholder="Insert the user and press enter"
                value={query}
                onChange={(e) => setQuery(e.target.value)}
                className="pl-3 border border-gray-300 rounded w-full lg:w-1/2 h-10"
                disabled={isLoading}
              />
            </div>
          </div>
        </form>
      </div>
    </div>
  );
};

export default SearchForm;