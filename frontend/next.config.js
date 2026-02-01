/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    domains: ['cdn.nba.com', 'a.espncdn.com'],
  },
}

module.exports = nextConfig
