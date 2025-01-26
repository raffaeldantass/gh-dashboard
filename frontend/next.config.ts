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
}

module.exports = nextConfig