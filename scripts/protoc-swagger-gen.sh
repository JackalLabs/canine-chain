#!/usr/bin/env bash

set -eo pipefail

mkdir -p ./tmp-swagger-gen
cd proto
proto_dirs=$(find canine_chain ../x -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf generate --template buf.gen.swagger.yaml $query_file
  fi
done

cd ..
# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./docs/config.json -o ./docs/swagger-ui/swagger.yaml -f yaml --includeDefinitions true

sed -E -i '' '/(items|properties): \*ref_[0-9]+/d' ./docs/swagger-ui/swagger.yaml
# clean swagger files
rm -rf ./tmp-swagger-gen