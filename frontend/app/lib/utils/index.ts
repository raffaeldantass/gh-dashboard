export const CookieUtils = {
  get: (name: string): string | null => {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    return parts.length === 2 ? parts.pop()?.split(';').shift() ?? null : null;
  },

  delete: (name: string): void => {
    document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/`;
  }
};

export const TokenManager = {
  get: (): string | null => {
    return sessionStorage.getItem('github_token') ?? CookieUtils.get('github_token');
  },

  set: (token: string): void => {
    sessionStorage.setItem('github_token', token);
  },

  clear: (): void => {
    sessionStorage.removeItem('github_token');
    CookieUtils.delete('github_token');
  }
};