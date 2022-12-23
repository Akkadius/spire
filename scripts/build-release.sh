#!/bin/bash

cwd=$(pwd)

# shellcheck disable=SC2046
if [ $(curl -s "https://api.github.com/repos/Akkadius/spire/tags" | jq -r '.[0].name' | sed 's/v//') = $(cat package.json | jq -r '.version') ]; then echo "Version tag is same as latest release exiting build"; exit; else echo "Local version different from remote, building..."; fi

# packr for packing web assets into binary
#go install github.com/gobuffalo/packr/packr

# This step originally packed static assets in the binary
# as of this writing it's gotten up to roughly 500-600MB uncompressed
#
# 2/23/2022 we are separating the assets from the binary and having them
# automatically downloaded separately on the initialization of the application
#
# Copy local asset images (This step is part of install now)
# make build-assets
# make strip-extra-assets

#:: Build SPA (Frontend)
cd "$cwd/frontend" && npm install && npm run build

#:: Pack frontend assets into binary
cd "$cwd" && make build-binary
cd "$cwd" && make release-binary
