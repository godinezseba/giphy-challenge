/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  async headers() {
    return [{
      source: '/:all*(gif|jpg|mp3|txt|png)',
      locale: false,
      headers: [{
        key: 'Cache-Control',
        value: 'public, max-age=172800, stale-while-revalidate=86400'
      }]
    }]
  }
}

module.exports = nextConfig
