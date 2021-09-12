START /wait taskkill /f /im spire.exe

go install github.com/gobuffalo/packr/packr

call npm install -g win-node-env

cd frontend && call npm run build & cd ..
xcopy "frontend\dist\" "public\" /s /e /y

packr clean
packr build