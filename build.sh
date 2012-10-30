#! /bin/sh

nex unql.nex 
goyacc unql.y
go build
