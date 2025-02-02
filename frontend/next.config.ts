/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'standalone',
  // Enable hot reload in Docker
  webpack: (config: { watchOptions: { poll: number; aggregateTimeout: number } }) => {
    config.watchOptions = {
      poll: 1000,
      aggregateTimeout: 300,
    }
    return config
  },
  async rewrites() {
    const isProduction = process.env.NODE_ENV === 'production';
    return [
      {
        source: '/api/:path*',
        destination: `http://${isProduction ? 'backend-prod' : 'backend-dev'}:8080/api/:path*`,
      },
    ];
  },
}

module.exports = nextConfig