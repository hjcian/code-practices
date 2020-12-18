#!/bin/bash

root=.
mkdir ${root}/logs

touch ${root}/logs/file1.log
echo -e "1477123675|ERROR|handler.cpp|127|findHandlers|Division by zero\n1477884474|TRACE|handler.cpp|130|findHandlers|Start calculations\n1477997395|ERROR|handler.cpp|1110|getHandlers|Array out of bounds\n1478944445|INFO|file.py|12|getData|Getting data\n1478975988|ERROR|do.sh|1700|start|errors\n" >> ${root}/logs/file1.log

touch ${root}/logs/file2.log
echo -e "1377123688|ERROR|handler.cpp|127|findHandlers|Division by zero\n1477123777|ERROR|hello.py|12|hell|Division by zero\n1477123558|ERROR|hello.py|127|findHandlers|Division by zero\n1477884474|TRACE|handler.cpp|130|findHandlers|Start calculations\n1477997395|ERROR|find-user.cpp|10|getUsers|Some error\n1478975945|INFO|file.py|22|writeData|Writing data\n1478975946|WARNING|file.scala|170|doSomething|just warning\n1477007000|ERROR|find-user.cpp|100|push|error\n" >> ${root}/logs/file2.log

touch ${root}/logs/file3.log
echo -e "1377123688|ERROR|handler.cpp|127|findHandlers|Division by zero\n1477123777|ERROR|hello.py|12|hell|Division by zero\n1477123558|ERROR|hello.py|127|findHandlers|Division by zero\n1477884474|TRACE|handler.cpp|130|findHandlers|Start calculations\n1477997395|ERROR|find-user.cpp|10|getUsers|Some error\n1478975945|INFO|file.py|22|writeData|Writing data\n1478975946|WARNING|ERROR|170|ERROR|ERROR warning\n1477007000|ERROR|find-user.cpp|100|push|error\n" >> ${root}/logs/file3.log

touch ${root}/logs/file.txt
echo -e "1377123688|ERROR|handler.cpp|127|findHandlers|Division by zero\n1477123777|ERROR|hello.py|12|hell|Division by zero\n1477123558|ERROR|hello.py|127|findHandlers|Division by zero\n1477884474|TRACE|handler.cpp|130|findHandlers|Start calculations\n1477997395|ERROR|find-user.cpp|10|getUsers|Some error\n1478975945|INFO|file.py|22|writeData|Writing data\n1478975946|WARNING|file.scala|170|doSomething|just warning\n1477007000|ERROR|find-user.cpp|100|push|error\n" >> ${root}/logs/file.txt

