#!/bin/sh

export ARCH=x86_64


if [ ! -f alpine-minirootfs-3.14.2-$ARCH.tar.gz ]; then
    wget https://dl-cdn.alpinelinux.org/alpine/v3.14/releases/x86_64/alpine-minirootfs-3.14.2-x86_64.tar.gz
fi

dd if=/dev/zero of=vda.img bs=1M count=512
mkfs.ext4 vda.img
sudo mkdir /mnt/vda
sudo mount vda.img /mnt/vda

prefix=$(pwd)
cd /mnt/vda
sudo tar xf ${prefix}/alpine-minirootfs-3.14.2-x86_64.tar.gz

sudo umount /mnt/vda
