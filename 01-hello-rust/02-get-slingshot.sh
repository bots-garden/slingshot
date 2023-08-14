#!/bin/bash
version="v0.0.0"
operating_system=$(uname -o) # Darwin GNU/Linux 
processor=$(uname -p) # x86_64 arm aarch64

if [ $operating_system == "GNU/Linux" ]; then
  operating_system="linux"
fi

if [ $operating_system == "Darwin" ]; then
  operating_system="darwin"
else
  operating_system="linux"
fi

if [ $processor == "x86_64" ]; then
  processor="amd64"
fi

if [ $processor == "arm" ]; then
  processor="arm64"
fi

if [ $processor == "aarch64" ]; then
  processor="arm64"
else
  processor="amd64"
fi

slingshot_target="slingshot-${version}-${operating_system}-${processor}"

wget https://github.com/bots-garden/slingshot/releases/download/${version}/${slingshot_target}

chmod +x ${slingshot_target}
mv ${slingshot_target} slingshot 
