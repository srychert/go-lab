#!/bin/bash

while getopts s: flag
do
    case "${flag}" in
        s) stop=${OPTARG};;
    esac
done

mkfifo p1 p2
./game > p1 < p2 &
./player -stop=$stop < p1 > p2

rm -rf p1 p2
