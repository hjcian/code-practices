#!/bin/bash
root=./files

mkdir -p ${root}
mkdir -p ${root}/dir1
mkdir -p ${root}/dir2

touch ${root}/dir1/file1.txt
echo -e "hello world 65.0.45.1 \nthis is some example 57.254.64.34\nfile with some correct and incorrect 54.76.4.56 ip addresses" >> ${root}/dir1/file1.txt
touch ${root}/dir1/file2.txt
echo -e "    hello world 87.0.65.76 \nthis is some example 36.5.76.35 \nfile with some correct and incorrect 15.128.4.65 ip addresses" >> ${root}/dir1/file2.txt
touch ${root}/dir2/file4.txt
echo -e "  file4.txt\n    hello world 127.98.0.1 \nthis is some example 128.96.107.55 \nfile with some correct and incorrect 128.68.4.11 ip addresses" >> ${root}/dir2/file4.txt
touch ${root}/dir2/file5.txt
echo -e "  hello world 127.0.49.1 \nthis is some example 128.99.58.55 \nfile with some correct and incorrect 26.56.4.23 ip addresses" >> ${root}/dir2/file5.txt
touch ${root}/dir2/file1.txt
echo -e "hello world 54.0.0.1 \nthis is some example 76.99.46.55 \nfile with some correct and incorrect 78.45.4.87 ip addresses\n" >> ${root}/dir2/file1.txt

touch ${root}/file1.txt
echo -e "hello world 65.0.45.1 \nthis is some example 57.254.64.34\nfile with some correct and incorrect 54.76.4.56 ip addresses" >> ${root}/file1.txt
touch ${root}/file2.txt
echo -e "hello world 0.0.0.0 \nthis is some example 57.2540.64.34\nfile with some correct and incorrect 54.76.400.56 ip addresses" >> ${root}/file2.txt
