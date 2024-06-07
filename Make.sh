#!/usr/bin/env bash
BuildSourcePackage="github.com/shaunfong/tabtoy/build"
BinaryPackage="github.com/shaunfong/tabtoy"
BinaryName="tabtoy"

BuildBinary()
{
  set -e
  TargetDir=bin/"${2}"
  mkdir -p "${TargetDir}"
  export GOOS=${2}
  BuildTime=$(date -R)
  GitCommit=$(git rev-parse HEAD)
  VersionString="-X \"${BuildSourcePackage}.BuildTime=${BuildTime}\" -X \"${BuildSourcePackage}.Version=${1}\" -X \"${BuildSourcePackage}.GitCommit=${GitCommit}\""

  go build -v -p 4 -o "${TargetDir}"/${BinaryName} -ldflags "${VersionString}" ${BinaryPackage}
  PackageDir=$(pwd)
  cd "${TargetDir}"
  tar zcvf "${PackageDir}"/${BinaryName}-${1}-"${2}"-x86_64.tar.gz ${BinaryName}
  cd "${PackageDir}"
}


if [[ ${2} == "" ]]; then
  BuildBinary windows
  BuildBinary linux
  BuildBinary darwin
else
  BuildBinary "${2}"
fi

