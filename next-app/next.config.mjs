/** @type {import('next').NextConfig} */
const nextConfig = {
    redirects: async () => {
        return [
            {
                source: '/',
                destination: '/products',
                permanent: true
            }
        ]
    },
    images: {
        remotePatterns: [
            {
                protocol: 'https',
                hostname: 'media.istockphoto.com'
            },
            {
                protocol: 'http',
                hostname: 'localhost'
            }
        ]
    }
};

export default nextConfig;
