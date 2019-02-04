#!/usr/bin/env bash

solcjs \
  --abi \
  --optimize \
  --output-dir ../solidity/build \
  --overwrite \
  ../solidity/contracts/Bridge.sol
