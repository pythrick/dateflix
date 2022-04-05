#!/bin/bash
# set -e
echo -e "\e[1;34m Start linter \e[0m"
golangci-lint --version
ROOT_DIR="${CI_PROJECT_DIR:-$(pwd)}"
echo -e "\e[1;36m Running golangci-lint on root folder \e[0m"
# regex to get linter msg and golangci-lint errors
golangci-lint run --allow-parallel-runners ./... --max-same-issues 0 | gawk '{ if(match($0, /([^.]+.go):([0-9]*):(.*)/, m) || match($0, /(^level=error)(.*)(msg.*)/, m)) print "* **" m[1] "** " m[2] ":" m[3]  }' 2>&1 > $ROOT_DIR/lintresult.tmp || true &
LS_DIR=$(find ./pkg -name "go.mod" | sed 's/go.mod//g')
for folder in $LS_DIR;
do
	echo -e "\e[1;36m Running golangci-lint on $folder \e[0m"
	cd $ROOT_DIR/$folder
	# regex to get linter msg and golangci-lint errors
  golangci-lint run --allow-parallel-runners ./... --max-same-issues 0 | gawk '{ if(match($0, /([^.]+.go):([0-9]*):(.*)/, m) || match($0, /(^level=error)(.*)(msg.*)/, m)) print "* **" FOLDER m[1] " ** " m[2] ":" m[3] }' FOLDER=$folder 2>&1 >> $ROOT_DIR/lintresult.tmp || true &
done
echo -e "\e[1;36m Processing files... \e[0m"
wait
cd $ROOT_DIR
results=$(cat lintresult.tmp)
if [ -z "$results" ]
then
	echo -e "\e[1;32mNo lint issues in your files! \e[0m"
else
  echo -e "\e[1;31m ***************************************** \e[0m"
  echo -e "\e[1;31m There are some lint issues in your files \e[0m"
  echo -e "\e[1;31m ***************************************** \e[0m"
  echo -e "\e[1;31m $results \e[0m"
  exit 1
fi
