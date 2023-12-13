PKGS=("cmd" "internal")
ROOT=$(pwd)
CODE=0

for pkg in "${PKGS[@]}"; do
  cd "$ROOT/$pkg" || exit 1
  echo -e "\n>>> Running linting on $pkg" ;

  # Run linting on the current package and its subpackages, excluding external packages
  GOGC=20 golangci-lint run --exclude ".*" "./..."

  # Check linting status
  if [[ "$?" != "0" ]]; then
    CODE=1
  fi
done

if [[ "$CODE" != 0 ]]; then
  exit 1
fi
