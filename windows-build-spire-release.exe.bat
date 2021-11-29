START /wait taskkill /f /im spire.exe

:: packr for packing web assets into binary
go install github.com/gobuffalo/packr/packr

:: Install node env
call npm install -g win-node-env

:: Build SPA (Frontend)
cd frontend && call npm run build & cd ..

:: Copy local asset images
:: curl -o %localappdata%\Temp\assets.zip -L https://github.com/Akkadius/eq-asset-preview/archive/refs/heads/master.zip
:: unzip -o %localappdata%\Temp\assets.zip -d %localappdata%\Temp\assets
:: xcopy "%localappdata%\Temp\assets\eq-asset-preview-master\assets\" "frontend\dist\assets\" /s /e /y

@REM xcopy "frontend\dist\" "public\" /s /e /y

:: Pack frontend assets into binary
packr clean
packr build
