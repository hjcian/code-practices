#!/bin/bash

prefix='./files'

mkdir -p ${prefix}/russia/moscow/center/b1
mkdir -p ${prefix}/russia/moscow/center/b2
mkdir -p ${prefix}/russia/moscow/east

touch ${prefix}/russia/moscow/center/b1/fan.info
touch ${prefix}/russia/moscow/center/b2/fan.info
touch ${prefix}/russia/moscow/east/fan.info

touch ${prefix}/invite.info
touch ${prefix}/ban.info

echo -n -e '00000001\n00000002\n00000003\n' >> ${prefix}/russia/moscow/center/b1/fan.info
echo -n -e '00000004\n00000005\n00000006\n' >> ${prefix}/russia/moscow/center/b2/fan.info
echo -n -e '10000007\n10000008\n10000009\n' >> ${prefix}/russia/moscow/east/fan.info

echo -n -e 'russia.moscow.center\nrussia.moscow.east\n' >> ${prefix}/invite.info
echo -n -e 'russia.moscow.center.b2\n' >> ${prefix}/ban.info
