export interface Repository {
  id: number;
  name: string;
  full_name: string;
  description: string | null;
  private: boolean;
  organization?: {
    login: string;
    avatar_url: string;
  };
  updated_at: string;
  language: string | null;
  total_pages?: number | null;
}

export interface SearchFilters {
  query: string;
  organization?: string;
  page?: number;
  per_page?: number;
}

export interface UseGithubReposReturn {
  repositories: Repository[];
  isLoading: boolean;
  error: string | null;
  currentPage: number;
  totalPages: number;
  handleSearchSubmit: (filters: Omit<SearchFilters, 'page' | 'per_page'>) => Promise<void>;
  handlePageChange: (page: number) => Promise<void>;
}

export interface PaginationProps {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
  isLoading: boolean;
}