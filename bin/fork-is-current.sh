#!/bin/bash
#
# check if the current fork has the latest commit from upstream
#

set -o errexit
set -o pipefail
set -o nounset

SCRIPT_HOME="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

#
# get the latest commit from upstream
#
UPSTREAM_LATEST=$(gh api repos/fileformat/ghashboard/commits | jq --raw-output '.[0].sha')
echo "INFO: upstream latest: $UPSTREAM_LATEST"

#
# GITHUB_REPOSITORY should be set when running in a GitHub Action
#
if [ "${GITHUB_REPOSITORY:-BAD}" == "BAD" ]; then
    echo "INFO: GITHUB_REPOSITORY not set, using local copy"
    #
    # see if it exists in the local repo
    #
    # git will return errlevel & script will exit if it isn't there
    # 
    # message will be:
    # fatal: ambiguous argument '...': unknown revision or path not in the working tree.
    #
    LOCAL_MESSAGE=$(git show -s --oneline 5${UPSTREAM_LATEST})
    echo "INFO: commit found in local repo!"
else
    echo "INFO: checking in ${GITHUB_REPOSITORY}..."
    # 
    # see if it exists in the current (forked) repo
    #
    # gh will return errlevel & script will exit if it isn't there
    #
    # message will be:
    # gh: No commit found for SHA: ... (HTTP 422)
    #
    FORKED_LATEST=$(gh api repos/${GITHUB_REPOSITORY}/commits/${UPSTREAM_LATEST})
    echo "INFO: commit found in current repo!"
fi