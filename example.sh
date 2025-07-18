#!/usr/bin/env nix
#!nix develop --command bash
# shellcheck shell=bash

go run cmd/invoice/main.go \
    generate --from "Dream, Inc." --to "Imagine, Inc." \
    --item "Rubber Duck" --quantity 1 --rate 25 \
    --tax 0.13 --discount 0.15 \
    --note "For debugging purposes."