#!/bin/sh

export ARCH=x86_64


if [ ! -f alpine-minirootfs-3.14.2-$ARCH.tar.gz ]; then
    wget https://dl-cdn.alpinelinux.org/alpine/v3.14/releases/x86_64/alpine-minirootfs-3.14.2-x86_64.tar.gz
fi

dd if=/dev/zero of=vda.img bs=1M count=512
mkfs.ext4 vda.img
mkdir /mnt/vda
sudo mount vda.img /mnt/vda

prefix=$(pwd)
cd /mnt/vda
sudo tar xf ${prefix}/alpine-minirootfs-3.14.2-x86_64.tar.gz
sudo rm -rf etc/conf.d etc/logrotate.d etc/modprobe.d etc/modules-load.d etc/network etc/opt etc/sysctl.d lib/modules-load.d
sudo rm etc/modules etc/udhcpd.conf etc/sysctl.conf etc/hostname etc/motd etc/issue etc/shadow
sudo cp /etc/passwd etc/passwd
sudo cp /etc/group etc/group
sudo cp /etc/hosts etc/hosts
sudo cp /etc/inittab etc/inittab
sudo cp /etc/fstab etc/fstab
sudo cp -Tr ../etc/init.d etc/init.d
sudo mkdir etc/dropbear
sudo cp -Tr /root root
sudo cp /usr/share/udhcpc/default.script usr/share/udhcpc/default.script
sudo rm -rf lib/sysctl.d
sudo rm -rf media opt srv

sudo umount /mnt/vda
