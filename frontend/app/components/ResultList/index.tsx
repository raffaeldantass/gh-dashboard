import React from 'react';
import Loading from '@/app/components/Loading';
import Pagination from '@/app/components/Pagination';
import { ResultListProps } from '@/app/types/github'

const ResultList: React.FC<ResultListProps> = ({ 
  repositories, 
  isLoading, 
  error, 
  currentPage,
  totalPages,
  onPageChange 
}) => {
  if (repositories.length === 0 && !isLoading) {
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
    <>
      <div className="mt-6 space-y-4">
        {isLoading ? (
          <Loading />
        ): (
          <>
            {repositories.map((repo) => (
              <div key={`${repo.owner}-${repo.name}-${repo.last_update.toString()}`} className="p-6 rounded-lg shadow-sm">
                <div className="flex items-start justify-between">
                  <div className="space-y-2">
                    <div className="flex items-center gap-2">
                      <h3 className="font-medium">{repo.name}</h3>
                    </div>
                    {repo.description ? (
                      <p className="text-sm text-gray-600">{repo.description}</p>
                    ) : ( <p className="text-sm text-gray-600">No description provided</p> )}
                    {repo.is_private && (
                      <p className="text-sm text-gray-600"> Private Repo </p>
                    )}
                  </div>
                  <div className="flex flex-col gap-2">
                    <span className="text-sm text-gray-500">
                    Last updated: {new Date(repo.last_update).toLocaleDateString('en-GB', {
                        day: '2-digit',
                        month: '2-digit',
                        year: 'numeric'
                      })}
                    </span>
                    <span className="text-sm text-gray-500">
                      Owner: {repo.owner}
                    </span>
                  </div>
                </div>
              </div>
            ))}
          </>
        )}
      </div>

      <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={onPageChange}
        isLoading={isLoading}
      />
    </>
  );
};

export default ResultList;