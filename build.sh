#! /bin/sh

cd parser/tuql
echo Running nex...
nex tuql.nex
echo Running goyacc...
goyacc tuql.y
cd ../..
echo Running go build...
go build
