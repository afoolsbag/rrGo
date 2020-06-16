#!/usr/bin/env bash

readonly SCRIPT_DIRECTORY="$(readlink --canonicalize-existing "$(dirname "${BASH_SOURCE[0]}")")"
readonly SCRIPT_DIRECTORY_NAME="$(basename "${SCRIPT_DIRECTORY}")"
readonly PROJECT_NAME="${SCRIPT_DIRECTORY_NAME/#docker-}"

command -v 'docker' &>'/dev/null' \
        || { echo 'The docker executable not found.'; \
             exit 1; }

test -e "${SCRIPT_DIRECTORY}/out/latest-build-version" \
        || { echo 'The latest-build-version not found.'; \
             exit 2; }
readonly LATEST_BUILD_VERSION="$(cat "${SCRIPT_DIRECTORY}/out/latest-build-version")"

docker image load --input "${SCRIPT_DIRECTORY}/out/${PROJECT_NAME}.${LATEST_BUILD_VERSION}.tgz" \
        || { echo 'docker load failed.'; \
             exit 3; }
echo 'docker load succeed.'

exit 0
