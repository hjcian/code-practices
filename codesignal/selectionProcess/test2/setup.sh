#!/bin/bash

# prefix = './root/devops'

# cd /
# rm -rf ./root/devops
# mkdir -p ./root/devops/
# cd ./root/devops/

mkdir -p school1

touch school1/maths.txt
touch school1/chemistry.txt
touch school1/geography.txt
touch school1/physics.txt
touch school1/informatics.txt
touch school1/biology.txt

echo -n -e 'A A 42\nA B 40\nC B 35\nB A 38\nD A 32\nF F 20\nG G 39\nH H 27\n' >> school1/maths.txt
echo -n -e 'CC CC 75\nEE EE 90\n' >> school1/chemistry.txt
echo -n -e 'Aaa Aab 39\nAaa Baa 37\nBaa Baa 38\n' >> school1/geography.txt
echo -n -e 'MMMM NNNN 12\n' >> school1/physics.txt
echo -n -e 'A A 250\nA B 275\n' >> school1/informatics.txt
echo -n -e 'B B 92\nT T 98\nL L 39\nK F 60\nS R 15\nT M 25\n' >> school1/biology.txt
