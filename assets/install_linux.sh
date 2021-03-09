#!/bin/bash
# Get route
path=$(pwd)
bashrc=/etc/bash.bashrc
# Add to $PATH
echo $PATH | grep -q "(^\|:\")$path\(:\|/\{0,1\}$\)" || echo "PATH=\$PATH:$path" >> "$bashrc"; source "$bashrc"