taskkill /im node.exe /f

SET NODE_OPTIONS=--openssl-legacy-provider node --max_old_space_size=4096 --stack-size=10000

cd frontend && npm install && node node_modules/@vue/cli-service/bin/vue-cli-service.js serve --open --host=0.0.0.0