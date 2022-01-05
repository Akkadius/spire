#!/bin/bash

cwd=$(pwd)

# packr for packing web assets into binary
go install github.com/gobuffalo/packr/packr

# Copy local asset images (This step is part of install now)
make build-assets

#:: Build SPA (Frontend)
cd "$cwd/frontend" && npm install && npm run build

#:: Pack frontend assets into binary
cd "$cwd" && make build-binary
cd "$cwd" && make release-binary
