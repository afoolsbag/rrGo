#!/usr/bin/env bash

readonly SCRIPT_DIRECTORY="$(readlink --canonicalize-existing "$(dirname "${BASH_SOURCE[0]}")")"
readonly SCRIPT_DIRECTORY_NAME="$(basename "${SCRIPT_DIRECTORY}")"
readonly PROJECT_NAME="${SCRIPT_DIRECTORY_NAME/#docker-}"

command -v 'docker' &>'/dev/null' \
        || { echo 'The docker executable not found.'; \
             exit 1; }

command -v 'gzip' &>'/dev/null' \
        || { echo 'The gzip executable not found.'; \
             exit 2; }

test -e "${SCRIPT_DIRECTORY}/out/latest-build-version" \
        || { echo 'The latest-build-version not found.'; \
             exit 3; }
readonly LATEST_BUILD_VERSION="$(cat "${SCRIPT_DIRECTORY}/out/latest-build-version")"

docker image save "${PROJECT_NAME}":"${LATEST_BUILD_VERSION}" | gzip > "${SCRIPT_DIRECTORY}/out/${PROJECT_NAME}.${LATEST_BUILD_VERSION}.tgz" \
        || { echo 'docker save failed.'; \
             exit 4; }
echo 'docker save succeed.'

exit 0
