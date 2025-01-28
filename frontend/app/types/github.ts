// types/github.ts

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
}

export interface SearchFilters {
  query: string;
  organization?: string;
}