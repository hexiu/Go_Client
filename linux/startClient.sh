#!/bin/bash
#
cd /mnt/
[  ! -d client ]&&mkdir client
cd client
nohub /mnt/client/main&
