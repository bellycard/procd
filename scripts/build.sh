#!/usr/bin/env bash

# Copyright 2013, Belly, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

# Aquire source directory, change to parent.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
PARENT_DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

cd $PARENT_DIR

# Specify extension for Windows build
EXTENSION=""
if [ "$(go env GOOS)" = "windows" ]; then
    EXTENSION=".exe"
fi

# Install dependencies
echo "[procd] Installing dependencies."
go get ./...

# Build Procd
echo "[procd] Building."
go build \
    -v \
    -o bin/procd${EXTENSION}
cp bin/procd${EXTENSION} $GOPATH/bin
