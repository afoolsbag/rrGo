#!/usr/bin/env bash

readonly SCRIPT_DIRECTORY="$(readlink --canonicalize-existing "$(dirname "${BASH_SOURCE[0]}")")"

command -v 'docker-compose' &>'/dev/null' \
        || { echo 'The docker-compose executable not found.'; \
             exit 1; }

cd "${SCRIPT_DIRECTORY}" \
        || { echo 'Change directory to script directory failed.'; \
             exit 2; }

docker-compose up --detach \
        || { echo 'docker-compose up and detach failed.'; \
             exit 3; }
echo 'docker-compose up and detach succeed.'

exit 0
