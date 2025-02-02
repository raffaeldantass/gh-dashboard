import { useGithubAuth } from "@/app/lib/hooks/useGithubAuth";

const GitHubAuth = () => {
  const { handleAuth, isAuthenticating, error } = useGithubAuth();

  return (
    <div className="mt-6 space-y-4 flex flex-col items-center justify-center w-full h-full">
      <p> Authenticate as a Github user to see Repos from Personal Account and Orgs. </p>
      <button 
        className="bg-blue-500 p-3 text-white rounded-md hover:bg-blue-600 transition-colors" 
        onClick={handleAuth}
      > 
         {isAuthenticating ? 'Authenticating...' : 'Login with GitHub'}
      </button>

      {error && (
        <div className="text-red-500">
          {error.message}
        </div>
      )}

    </div>
  );
}

export default GitHubAuth;