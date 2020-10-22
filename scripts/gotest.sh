#!/usr/bin/env bash
# scripts to help for run go test case
# ./gotest.sh <package name> <test case> <package name> ...

#find ./* -type f -name "*.go" -exec dirname {} \; | sort|uniq |xargs -i go test -v {} -test.run=".*";

CURRENT_PATH=$(cd "$(dirname "$0")" || exit 1; pwd)
cd "${CURRENT_PATH}"/.. || exit 1

ALL="."
CASES=()

if [[ "$#" == "0" ]]; then
  # exec all test case
  while IFS='' read -r line; do
    CASES+=("${line}:${ALL}");
  done < <(find ./ -type f -name '*.go' -exec dirname {} \; |sort |uniq |xargs -n1 -I X echo X |sed 's@^./@@g')
else
  # exec specifeid test case
  args=("$@")
  for arg in "${args[@]}";do
    arg="$(echo "$arg"|sed 's@^./@@g; s@/$@@g')"

    if echo "$arg"|grep -q '^Test'; then
      while IFS='' read -r line; do
          CASES+=("${line}:${arg}")
      done < <(grep -rl "^func ${arg}(" |xargs -n1 -I X dirname X)
    else
      if test -d "$arg"; then
        CASES+=("${arg}:${ALL}")
      else
        echo "Error: package <${arg}> not exists"
        exit 1
      fi
    fi
  done
fi

for CASE in "${CASES[@]}";do
  pair=(${CASE//:/ })

  p="${pair[0]}"
  c="${pair[1]}"

  go test -race -v "./${p}" -test.run="${c}" -test.count=1 -failfast || exit 1
done

exit 0
