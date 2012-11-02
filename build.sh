#! /bin/sh

cd parser
echo Running nex...
nex unql.nex 
echo Running goyacc...
goyacc unql.y
cd ..
echo Running go build...
go build
