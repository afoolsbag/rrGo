#!/usr/bin/env bash

readonly SCRIPT_DIRECTORY="$(readlink --canonicalize-existing "$(dirname "${BASH_SOURCE[0]}")")"
readonly SCRIPT_DIRECTORY_NAME="$(basename "${SCRIPT_DIRECTORY}")"
readonly PROJECT_NAME="${SCRIPT_DIRECTORY_NAME/#docker-}"
readonly BUILD_VERSION="$(date +'%y%m%d.%H%M')"

command -v 'docker' &>'/dev/null' \
        || { echo 'The docker executable not found.'; \
             exit 1; }

cd "${SCRIPT_DIRECTORY}/src" \
        || { echo 'Change directory to source directory failed.'; \
             exit 2; }

docker builder build --tag "${PROJECT_NAME}":'latest' --tag "${PROJECT_NAME}":"${BUILD_VERSION}" '.' \
        || { echo 'docker build failed.'; \
             exit 3; }
echo 'docker build succeed.'

mkdir --parents "${SCRIPT_DIRECTORY}/out" \
        || { echo 'Make out directory failed.'; \
             exit 4; }

echo "${BUILD_VERSION}" | tee "${SCRIPT_DIRECTORY}/out/latest-build-version" \
        || { echo 'Write latest-build-version failed.'; \
             exit 5; }

exit 0
