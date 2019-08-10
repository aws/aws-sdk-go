#!/bin/sh

HEAD_PATH=`pwd`
SRC=$1
OUT=$2
if [ -z "${OUT}" ]; then
	echo "output path not specified"
	exit 1
fi

cd ${SRC}

rm -rf ${OUT}
mkdir -p ${OUT}

for PKG_PATH in $(find ./aws ./service -type d)
do
	if [ -z "$(ls $PKG_PATH/*.go 2> /dev/null)" ]; then
		continue
	fi
	OUTFILE=${OUT}/${PKG_PATH//[\.\/]/_}.apidiff
	$(go env GOPATH)/bin/apidiff -w ${OUTFILE} ${PKG_PATH}
done
