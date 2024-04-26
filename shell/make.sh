#!/bin/bash

LOGO=$(cat "$PWD"/shell/logo.txt)
GITHOOKS_PATH="$PWD"/.githooks
GIT_HOOKS_PATH="$PWD"/.git/hooks

source "$PWD"/shell/colors.sh
source "$PWD"/shell/ascii.sh

echo "${bold}${blue}$LOGO${normal}"

setupGitHooks() {
  title1 "Setting up git hooks..."

  find $GITHOOKS_PATH -type f -exec cp {} $GIT_HOOKS_PATH \;
  chmod +x .git/hooks/*

  lineOk "All hooks installed and updated"
}

checkHooks() {
  err=0
  echo "Checking git hooks..."
  for FILE in "$GITHOOKS_PATH"/*; do
    f="$(basename -- $FILE)"
    FILE2="$GIT_HOOKS_PATH"/$f
    if [ -f "$FILE2" ]; then
      if cmp -s "$FILE" "$FILE2"; then
        lineOk "Hook file ${underline}$f${normal} installed and updated"
      else
        lineError "Hook file ${underline}$f${normal} ${red}installed but out-of-date [OUT-OF-DATE]"
        err=1
      fi
    else
      lineError "Hook file ${underline}$f${normal} ${red}not installed [NOT INSTALLED]"
      err=1
    fi
    if [ $err -ne 0 ]; then
      echo -e "\nRun ${bold}make setup-git-hooks${normal} to setup your development environment, then try again.\n"
      exit 1
    fi
  done

  lineOk "\nAll good"
}

lint() {
  title1 "STARTING LINT"
  out=$(golangci-lint run --fix ./... 2>&1)
  out_err=$?
  perf_out=$(perfsprint ./... 2>&1)
  perf_err=$?

  echo "$out"
  echo "$perf_out"

  if [ $out_err -ne 0 ]; then
    echo -e "\n${bold}${red}An error has occurred during the lint process: \n $out\n"
    exit 1
  fi
  if [ $perf_err -ne 0 ]; then
    echo -e "\n${bold}${red}An error has occurred during the performance check: \n $perf_out\n"
    exit 1
  fi

  lineOk "Lint and performance checks passed successfully"
}

format() {
  title1 "Formatting all golang source code"
  gofmt -w ./
  lineOk "All go files formatted"
}

case "$1" in
"setupGitHooks")
  setupGitHooks
  ;;
"checkHooks")
  checkHooks
  ;;
"lint")
  lint
  ;;
"format")
  format
  ;;
*)
  echo -e "\n\n\nExecuting with parameter $1"
  makeCmd "$1"
  ;;
esac
