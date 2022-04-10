#!/bin/bash

mkfifo p1 p2
./game > p1 < p2 &
./player < p1 > p2

rm -rf p1 p2
