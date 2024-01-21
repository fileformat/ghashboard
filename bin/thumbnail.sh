#!/bin/bash
#
# make thumbnails of the screenshots
#

set -o errexit
set -o pipefail
set -o nounset

SCRIPT_HOME="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
IMAGE_HOME=$(realpath "${SCRIPT_HOME}/../docs/images")

SCREENSHOTS=( "08_result.png" "10_customized_result.png" )

for SCREENSHOT in "${SCREENSHOTS[@]}"
do
  echo "Processing ${SCREENSHOT}"
  # -resize 25%
  convert "${IMAGE_HOME}/screenshots/${SCREENSHOT}" -thumbnail 320x -unsharp 0x.5 "${IMAGE_HOME}/thumbnails/${SCREENSHOT}"
done