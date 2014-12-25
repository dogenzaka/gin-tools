#!/bin/bash

echo "mode: count" > coverage.txt

Dirs=("./ ")
Dirs+=$(find ./* -maxdepth 3 -not -path '*/.git*' -not -path '*/.*' -not -path '*/_*' -type d )
for Dir in $Dirs;
do
        if ls $Dir/*.go &> /dev/null;
        then
            go test -v -covermode=count -coverprofile=temp.out $Dir
            if [ -f temp.out ]
            then
                cat temp.out | grep -v "mode: count" >> coverage.txt
            fi
fi
done

rm temp.out
