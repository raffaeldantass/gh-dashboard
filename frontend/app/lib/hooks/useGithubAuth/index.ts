import { useCallback, useState } from 'react';
import { fetchApi } from '@/app/lib/api';
import { AuthError, AuthResponse, UseGithubAuthReturn } from '@/app/types/github';

export const useGithubAuth = (): UseGithubAuthReturn => {
  const [isAuthenticating, setIsAuthenticating] = useState(false);
  const [error, setError] = useState<AuthError | null>(null);

  const handleAuth = useCallback(async () => {
    setIsAuthenticating(true);
    setError(null);

    try {
      const response = await fetchApi('/login');
      const data = response as AuthResponse;

      if (!data?.url) {
        throw new Error('Invalid authentication response');
      }

      // Redirect to GitHub OAuth flow
      window.location.href = data.url;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to authenticate with GitHub';
      
      setError({
        message: errorMessage,
        code: 'AUTH_ERROR'
      });
      
      console.error('GitHub authentication error:', err);
    } finally {
      setIsAuthenticating(false);
    }
  }, []);

  return {
    handleAuth,
    isAuthenticating,
    error
  };
};