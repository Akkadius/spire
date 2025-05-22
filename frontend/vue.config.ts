import {defineConfig} from 'vite'

export default defineConfig(({}) => {
    return {
        transpileDependencies: true,
        build: {
            rollupOptions: {
                // eslint-disable-next-line no-undef
                external: new Regexp('public/eq-asset-preview-master/.*')
            }, // ...etc.
        }
    }
})
