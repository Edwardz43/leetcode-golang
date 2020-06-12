#!/bin/sh
FILENAME=""
for i in "${@}"
do
    if [[ $i =~ ^[0-9]+.$ ]]; then
        FILENAME="${FILENAME}${i//./}"
    else
        FILENAME="${FILENAME}_${i}"
    fi
done

mkdir ${FILENAME} && mkdir ${FILENAME}/solution
cp ./solution_template ${FILENAME}/solution/solution.go
cp ./solution_test_template ${FILENAME}/solution/solution_test.go