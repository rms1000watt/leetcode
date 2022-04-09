#!/usr/bin/env bash

set -e

if [[ -z ${1} ]]; then
  echo "ERROR: \${1} tag version is empty"
  exit 1
fi

git add .
git commit -m "running tag-push.sh"
git push origin master
git tag "${1}"
git push origin "${1}"
