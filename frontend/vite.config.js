import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

/**
 * @type {import('vite').UserConfig}
 */
export default {
    plugins: [vue()],
    alias: {
        '@components': resolve(__dirname, 'src/components'),
        '@api': resolve(__dirname, 'src/api'),
    },
    server: {
        proxy: {
            // target is 'backend' because ...
            '^/api/': {
                target: 'http://backend:8000',
                changeOrigin: true,
                secure: false,
            },
        },
    },
}
