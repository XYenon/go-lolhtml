#!/usr/bin/env bash

set -xeuo pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

cp "$DIR/src/c-api/include/lol_html.h" "$DIR/include/"

declare -A targets
targets["x86_64-unknown-linux-musl"]="linux-x86_64"
targets["aarch64-unknown-linux-musl"]="linux-aarch64"
targets["x86_64-apple-darwin"]="macos-x86_64"
targets["aarch64-apple-darwin"]="macos-aarch64"
for target in "${!targets[@]}"; do
  docker run --rm \
      --volume "$DIR/src":/root/src \
      --workdir /root/src/c-api \
        joseluisq/rust-linux-darwin-builder:latest \
          sh -c "cargo build --release --target $target"
  mkdir -p "$DIR/${targets[$target]}"
  cp "$DIR/src/c-api/target/$target/release/liblolhtml.a" "$DIR/${targets[$target]}/"
done
