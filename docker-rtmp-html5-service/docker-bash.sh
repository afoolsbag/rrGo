#!/usr/bin/env bash

readonly SCRIPT_DIRECTORY="$(readlink --canonicalize-existing "$(dirname "${BASH_SOURCE[0]}")")"
readonly SCRIPT_DIRECTORY_NAME="$(basename "${SCRIPT_DIRECTORY}")"
readonly PROJECT_NAME="${SCRIPT_DIRECTORY_NAME/#docker-}"

command -v 'docker' &>'/dev/null' \
        || { echo 'The docker executable not found.'; \
             exit 1; }

docker container run --net host --interactive --rm --tty "${PROJECT_NAME}" bash \
        || { echo 'docker run failed.'; \
             exit 2; }

exit 0
