export interface Repository {
  id: number;
  name: string;
  description?: string | null;
  is_private: boolean;
  owner: string;
  last_update: string;
}

export interface ResultListProps {
  repositories: Repository[];
  isLoading: boolean;
  error?: string | null;
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
}

export interface PaginatedResponse {
  repositories: Repository[];
  current_page: number;
  total_pages: number;
  total_items: number;
  items_per_page: number;
}

export interface UseGithubReposReturn {
  repositories: Repository[];
  isLoading: boolean;
  error: string | null;
  currentPage: number;
  totalPages: number;
  fetchRepositories: (page: number) => Promise<void>;
  handlePageChange: (page: number) => Promise<void>;
  logout: () => void;
}

export interface AuthResponse {
  url: string;
}

export interface AuthError {
  message: string;
  code?: string;
}

export interface UseGithubAuthReturn {
  handleAuth: () => Promise<void>;
  isAuthenticating: boolean;
  error: AuthError | null;
}