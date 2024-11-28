#!/bin/sh

src_files=$(find . -type f \( -iname \*.cpp -o -iname \*.h \))

for file in $src_files
do
    clang-format $file > .tmp && mv .tmp $file
done
