#!/usr/bin/env bash

set -Eeuo pipefail

readonly REPO_URL="https://github.com/meleu/gordm"

readonly SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"

generate_readme() {
  cat "${SCRIPT_DIR}/readme-header.md"
  commit2md
}

# Print all commit messages as markdown, oldest first
commit2md() {
  local commit_base_url="${REPO_URL}/commit"
  local title hash body

  git log --reverse --pretty=format:'%s%x00%H%x00%b%x00' \
    | while IFS= read -r -d '' title \
      && IFS= read -r -d '' hash \
      && IFS= read -r -d '' body; do

      [[ "$title" == *README.md* || "$title" == *"First commit"* ]] && continue

      echo "### ${title//$'\n'/}"
      echo -e "\n[link to commit](${commit_base_url}/${hash})\n"
      echo "$body"
    done
}

generate_readme
