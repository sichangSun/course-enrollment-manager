import { fileURLToPath, URL } from 'node:url'

import { defineConfig,loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'
import VueDevTools from 'vite-plugin-vue-devtools'

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
	const env = loadEnv(mode, process.cwd(), '')

  return {
    // vite config
    define: {
      __APP_ENV__: JSON.stringify(env.APP_ENV),
			_BASE_URL_: JSON.stringify(env.VITE_BASE_URL)

    },
		plugins: [
			vue(),
			VueDevTools(),
		],
		resolve: {
			alias: {
				'@': fileURLToPath(new URL('./src', import.meta.url)),
			},
		},
		server: {
			host: true,
		}
  }

})
