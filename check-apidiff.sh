#!/bin/sh

OLD=$1
NEW=$2

for PKG_META in $(find ${OLD} -name "*.apidiff" | sed -e "s|^${OLD}||")
do
	$(go env GOPATH)/bin/apidiff ${OLD}${PKG_META} ${NEW}${PKG_META}
done

