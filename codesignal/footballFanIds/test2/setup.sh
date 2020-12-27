#!/bin/bash

prefix='./files'

mkdir -p ${prefix}/russia/city/moscow/group/center/b1/a1

touch ${prefix}/russia/city/moscow/group/center/b1/a1/fan.info

touch ${prefix}/invite.info
touch ${prefix}/ban.info

echo -n -e '00000001\n00000002\n00000003\n' >> ${prefix}/russia/city/moscow/group/center/b1/a1/fan.info

echo -n -e 'russia\nrussia.city.moscow\nrussia.city.moscow.group.center\nrussia.city.moscow.group.center.b1.a1\n' >> ${prefix}/invite.info
echo -n -e 'russia.city\nrussia.city.moscow.group\nrussia.city.moscow.group.center.b1\n' >> ${prefix}/ban.info
