#!/bin/sh
#
# This script is *heavily* influenced by the wonderful people
# on CoreOS' etcd team. Take a look at their test script here:
#
# https://github.com/coreos/etcd
#
# They are pretty much the original authors.


# all github.com/fschr/go/whatever pkgs that are not auto-generated / tools
PKGS=`find . -name \*.go | while read a; do dirname $a; done | sort | uniq | sed "s|\.|${REPO_PATH}|g"`
# pkg1,pkg2,pkg3
PKGS_COMMA=`echo ${PKGS} | sed 's/ /,/g'`

TEST_PKGS=`find . -name \*_test.go | while read a; do dirname $a; done | sort | uniq | sed "s|\./||g"`
FORMATTABLE=`find . -name \*.go | while read a; do echo $(dirname $a)/"*.go"; done | sort | uniq | sed "s|\./||g"`

if [ -z "$FORMATTABLE" ]; then
    echo -e "no formattable *.go files found"
    exit 255
fi

if [ -z "$PASSES" ]; then
	PASSES="build fmt" # space separated list of functions
fi

build_pass () {
    echo "Checking 'go install'..."
    installRes=$(go install)
    if [ -n "${installRes}" ]; then
	echo -e "'go install' failed:\n${installRes}"
	exit 255
    fi
}

fmt_pass () {
	echo "Checking gofmt..."
	fmtRes=$(gofmt -l -s -d $FORMATTABLE)
	if [ -n "${fmtRes}" ]; then
		echo -e "gofmt checking failed:\n${fmtRes}"
		exit 255
	fi

	echo "Checking govet..."
	if ! [ -z "$TEST_PKGS" ]; then
	    vetRes=$(go vet $TEST_PKGS)
	    if [ -n "${vetRes}" ]; then
		echo -e "govet checking failed:\n${vetRes}"
		exit 255
	    fi
	fi

	echo "Checking 'go tool vet -shadow'..."
	for path in $FORMATTABLE; do
		if [ "${path##*.}" != "go" ]; then
			path="${path}/*.go"
		fi
		vetRes=$(go tool vet -shadow ${path} 2>&1)
		if [ -n "${vetRes}" ]; then
			echo -e "govet -shadow checking ${path} failed:\n${vetRes}"
			exit 255
		fi
	done
}

for pass in $PASSES; do
	${pass}_pass $@
done
