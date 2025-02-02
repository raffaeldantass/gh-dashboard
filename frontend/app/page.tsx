'use client';

import GitHubAuth from '@/app/components/GithubAuth';

export default function Home() {
  return (
    <main className="container mx-auto px-4 py-8 h-screen">
      <h1 className="text-3xl font-bold mb-8">GitHub Repositories</h1>
      <GitHubAuth />
    </main>
  );
}