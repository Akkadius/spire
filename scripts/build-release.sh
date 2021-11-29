# packr for packing web assets into binary
go install github.com/gobuffalo/packr/packr

#:: Build SPA (Frontend)
cd frontend && npm run build && cd ..

# Copy local asset images (This step is part of install now)
# curl -o /tmp/assets.zip -L https://github.com/Akkadius/eq-asset-preview/archive/refs/heads/master.zip
# unzip -o /tmp/assets.zip -d /tmp/assets
# cp -R /tmp/assets/eq-asset-preview-master/assets/ ./frontend/dist/

#:: Pack frontend assets into binary
packr clean
packr

go build
GOOS=windows GOARCH=amd64 go build

gh-release --assets=./spire -y
