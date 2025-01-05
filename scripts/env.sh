#!/bin/bash
set -e

echo "GPG_PRIVATE_KEY=${GPG_PRIVATE_KEY}" > .env
echo "GPG_PASSPHRASE=${GPG_PASSPHRASE}" >> .env
echo "GPG_PUBLIC_KEY=${GPG_PUBLIC_KEY}" >> .env
echo "ARCH=${ARCH}" >> .env
