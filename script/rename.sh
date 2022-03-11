#!/usr/bin/env bash

cd $1
find . -type f -iname "*.mod*" -exec sed -i '' -e "s%$1%$2%g" {} \;
find . -type f -iname "*.go*" -exec sed -i '' -e "s%$1%$2%g" {} \;
