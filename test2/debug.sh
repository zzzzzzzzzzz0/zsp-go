#!/bin/bash
d=`dirname $0`
cd $d/root
$PWD/../../bin/zsp -zhscript-o-tree -zhscript-o-path -zhscript-o-lc -zhscript-o-ansi ---- -a :4001 -r $PWD/
