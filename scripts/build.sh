#!/bin/sh

VERSION="v0.0.0"
SHA=`git rev-parse HEAD 2> /dev/null || true`
DIRTY=`git status --porcelain --untracked-files=no`
PACKAGE="github.com/bsatlas/aws-rget/cmd/version"

if [ "$DIRTY" ]
then
	COMMIT="$SHA-dirty"
else
	COMMIT="$SHA"
fi

go build \
-ldflags "-X $PACKAGE.Version=$VERSION -X $PACKAGE.Commit=$COMMIT" \
-o bin/aws-rget ./cmd

