#!/bin/bash

# prefix = './root/devops'

# cd /
# rm -rf ./root/devops
# mkdir -p ./root/devops/
# cd ./root/devops/

mkdir -p school1
mkdir -p school2
mkdir -p school3

touch school1/maths.txt
touch school1/chemistry.txt
touch school2/maths.txt
touch school3/maths.txt
touch school3/physics.txt
touch school3/chemistry.txt

echo -n -e 'James Davis 42\nDaniel Smith 40\nJessica Robinson 35\nDavid Thompson 38\n' >> school1/maths.txt
echo -n -e 'Charles Clark 75\nAnthony Lee 90\n' >> school1/chemistry.txt
echo -n -e 'Kenneth Wilson 39\nRobert Brown 37\n' >> school2/maths.txt
echo -n -e 'Michael Phillips 41\n' >> school3/maths.txt
echo -n -e 'Harry Nelson 55\nLouis Gonzalez 66\n' >> school3/physics.txt
echo -n -e 'Diana Wood 80\nRalph Jordan 100\nJuan Gomez 57\nBruce Peterson 62\nAlan Kelly 38' >> school3/chemistry.txt
