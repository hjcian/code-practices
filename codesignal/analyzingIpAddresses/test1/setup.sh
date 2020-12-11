#!/bin/bash
root=./files

mkdir -p ${root}
mkdir -p ${root}/dir1
mkdir -p ${root}/dir2

touch ${root}/dir1/file1.txt
echo -e "hello world 127.0.0.1\nthis is some example 128.99.107.55 \nfile with some correct and incorrect 128.128.4.11 ip 0.11.1115.78 addressesaddresses" >> ${root}/dir1/file1.txt
touch ${root}/dir1/file2.txt
echo -e "hello from 74.0.65.76 and 8.dd.99.88.907 good\nthis is some example 16.1215.76.35 \nfile with some correct and incorrect 15.128.4.65 ip addresses\n0.0.0.0" >> ${root}/dir1/file2.txt
touch ${root}/dir2/file3.txt
echo -e "127.65.64.1 127.0.64.1 127.0.0.1\nexample 128.57.107.76 128.57.907.70 \nfile with some correct and incorrect 67.128.4.11 ip addresses 7.7.7.8" >> ${root}/dir2/file3.txt
touch ${root}/dir2/file4.txt
echo -e "hello world 127.98.0.1\nthis is some example 128.96.107.55 \nfile with some correct and incorrect 128.68.4.11 ip addresses" >> ${root}/dir2/file4.txt
touch ${root}/f.inp
echo -e "hello world 127.0.49.1 \nthis is some example 128.99.58.55 8.88.888.88 77.255.255.254\n7.7.257.25 file with some correct and incorrect 26.56.4.23 ip addresses" >> ${root}/f.inp