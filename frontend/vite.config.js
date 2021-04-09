import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

export default {
    plugins: [vue()],
    alias: {
        '@': resolve(__dirname, './src')
    },
    server: {
        proxy: {
            // target is 'backend' because docker-compose resolves host by service name
            // https://stackoverflow.com/questions/61823628/getting-vue-devserver-proxy-to-work-with-different-local-ports-in-docker
            '^/api/': {
                target: 'http://backend:8000',
                changeOrigin: true,
                secure: false,
            },
        },
    },
}
