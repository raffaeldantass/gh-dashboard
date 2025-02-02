export const CookieUtils = {
  get: (name: string): string | null => {
    if (typeof window === 'undefined') {
      return null;
    }
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    return parts.length === 2 ? parts.pop()?.split(';').shift() ?? null : null;
  },

  delete: (name: string): void => {
    if (typeof window !== 'undefined') {
      document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/`;
    }
  }
};

export const TokenManager = {
  get: (): string | null => {
    try {
      return sessionStorage.getItem('github_token') ?? CookieUtils.get('github_token');
    } catch {
      return null;
    }
  },

  set: (token: string): void => {
    try {
      sessionStorage.setItem('github_token', token);
    } catch {
      // Handle error silently
    }
  },

  clear: (): void => {
    try {
      sessionStorage.removeItem('github_token');
      CookieUtils.delete('github_token');
    } catch {
      // Handle error silently
    }
  }
};