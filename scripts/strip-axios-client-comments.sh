#!/bin/bash

find ./frontend/src/app/api/api/ -type f -name \*.ts | while read f; do cat $f | sed -n '/^\/\*.*\*\//!p' | sed -n '/ \/\/.*/!p' | sed 's|/\*|\n&|g;s|*/|&\n|g' | sed '/\/\*/,/*\//d' | sed '/^$/d' | sed '/^[[:space:]]*$/d' | tee $f.bak; done

find ./frontend/src/app/api/api/ -name "*.bak" -exec sh -c 'mv -f $0 ${0%.bak}' {} \;
