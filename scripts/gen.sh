#!/usr/bin/env bash

d=$1

if [[ -z $d ]];
then
  echo "error Please provide day (d) parameter, e.g., make gen 1"
  exit 1
fi

# Create the folder if it doesn't exist
mkdir -p "day$d";
# Create an empty file inside the folder
echo "package day$d

import (
	\"fmt\"
)

const INPUTS = \"./day$d/input.txt\"

type Challenge struct{}

func (d Challenge) A() {
	fmt.Println(\"function A() not yet implemented\")
}

func (d Challenge) B() {
	fmt.Println(\"function B() not yet implemented\")
}
" > "day$d/day$d.go";
touch "day$d/input.txt";
