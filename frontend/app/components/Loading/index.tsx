const Loading: React.FC = () => (
  <div className="space-y-4">
    {[1, 2, 3].map((i) => (
      <div key={i} className="p-6 rounded-lg shadow-sm animate-pulse">
        <div className="space-y-3">
          <div className="h-4 bg-gray-200 rounded w-[250px]" />
          <div className="h-4 bg-gray-200 rounded w-[200px]" />
          <div className="h-4 bg-gray-200 rounded w-[150px]" />
        </div>
      </div>
    ))}
  </div>
);

export default Loading;
