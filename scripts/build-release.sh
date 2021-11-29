#!/bin/bash

cwd=$(pwd)

# packr for packing web assets into binary
go install github.com/gobuffalo/packr/packr

# Copy local asset images (This step is part of install now)
curl --compressed -o /tmp/assets.zip -L https://github.com/Akkadius/eq-asset-preview/archive/refs/heads/master.zip
unzip -o /tmp/assets.zip -d /tmp/assets
cp -R /tmp/assets/eq-asset-preview-master/ ./frontend/public/

#:: Build SPA (Frontend)
cd "$cwd/frontend" && npm install && npm run build

#:: Pack frontend assets into binary
cd "$cwd" && packr clean
cd "$cwd" && packr
cd "$cwd" && go build
cd "$cwd" && GOOS=windows GOARCH=amd64 go build
cd "$cwd" && gh-release --assets=./spire./spire.exe -y
