import React from 'react';
import Image from 'next/image';
import type { Repository } from '@/app/types/github';
import Loading from '@/app/components/Loading';

interface ResultListProps {
  repositories: Repository[];
  isLoading: boolean;
  error?: string | null;
}

const ResultList: React.FC<ResultListProps> = ({ repositories, isLoading, error }) => {
  if (repositories.length === 0) {
    return (
      <>
        {error ? (
          <div className="bg-red-50 text-red-500 p-4 rounded-lg mb-4">
            {error}
          </div>
        ) : (
          <div className="p-6 rounded-lg text-center text-gray-500">
            No data to display.
          </div>
        )}
      </>
    );
  }

  return (
    <div className="mt-6 space-y-4">
      {isLoading ? (
        <>
        <Loading />
        <h1>Here</h1>
        </>
      ): (
        <>
          {repositories.map((repo) => (
            <div key={repo.id} className="p-6 rounded-lg shadow-sm">
              <div className="flex items-start justify-between">
                <div className="space-y-2">
                  <div className="flex items-center gap-2">
                    {repo.private ? (
                      <div className="h-4 w-4 text-amber-500" ></div>
                    ) : (
                      <div className="h-4 w-4 text-green-500" ></div>
                    )}
                    <h3 className="font-medium">{repo.name}</h3>
                  </div>
                  {repo.description && (
                    <p className="text-sm text-gray-600">{repo.description}</p>
                  )}
                  <div className="flex items-center gap-4 text-sm text-gray-500">
                    {repo.organization && (
                        <Image
                          src={repo.organization.avatar_url}
                          alt={repo.organization.login}
                          width={16}
                          height={16}
                          className="rounded-full"
                        />
                    )}
                    {repo.language && (
                      <div className="flex items-center gap-1">
                        <div className="h-4 w-4" ></div>
                        <span>{repo.language}</span>
                      </div>
                    )}
                  </div>
                </div>
                <span className="text-sm text-gray-500">
                  Updated {new Date(repo.updated_at).toLocaleDateString()}
                </span>
              </div>
            </div>
          ))}
        </>
      )}
    </div>
  );
};

export default ResultList;