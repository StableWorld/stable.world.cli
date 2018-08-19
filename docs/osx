#!/bin/sh
set -eou pipefail

LOG="$TMPDIR/sw-installer.log"

logecho(){
  echo "$*"
  echo "$*" > $LOG
}
err_report() {
    echo "Error on line $1. Check '$LOG' for details"
}

trap 'err_report $LINENO' ERR

POSITIONAL=()
VERSION=""
INSTALL_PREFIX="/usr/local/bin"

while [[ $# -gt 0 ]]; do
  key="$1"
  case $key in
      -v|--version)
      VERSION="$2"
      shift # past argument
      shift # past value
      ;;
      -p|--prefix)
      INSTALL_PREFIX="$2"
      shift # past argument
      shift # past value
      ;;
      *)    # unknown option
      POSITIONAL+=("$1") # save it in an array for later
      shift # past argument
      ;;
  esac
done
set -- "${POSITIONAL[@]}" # restore positional parameters

logecho "VERSION: $VERSION"
logecho "INSTALL_PREFIX: $INSTALL_PREFIX"
if [ -z "$VERSION" ]; then
  logecho "Fetch Latest Version ... "
  logecho "https://api.github.com/repos/$REPO/releases/latest"
  VERSION=$(curl  https://api.github.com/repos/$REPO/releases/latest 2> "$LOG" | grep tag_name | sed 's/.*"\(v.*\)".*/\1/g')
fi
logecho "Using version $VERSION"

REPO="StableWorld/stable.world.cli"
OS="darwin"
ARCH="amd64"
COMMAND=$1

logecho "Downloading '$COMMAND'"
COMMAND_PATH="$TMPDIR/$COMMAND"
logecho "https://github.com/$REPO/releases/download/$VERSION/$COMMAND-$OS-$ARCH"
curl -o "$COMMAND_PATH" -L --fail "https://github.com/$REPO/releases/download/$VERSION/$COMMAND-$OS-$ARCH" 2> "$LOG"

chmod +x "$COMMAND_PATH"
mv "$COMMAND_PATH" "$INSTALL_PREFIX/$COMMAND"
# https://github.com/StableWorld/stable.world.cli/releases/download/v0.0.4/scurl-linux-amd64
