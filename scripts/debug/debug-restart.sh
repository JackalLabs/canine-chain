#!/bin/bash

KEY="j1"
KEY1="j2"
KEY2="charlie"
DEPOACCKEY="deposit_account"

CHAINID="test-1"
MONIKER="localjack"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"
BROADCASTMODE="block"

# Start the node with debugger
screen -d -m -S "canined" bash -c "canined start --pruning=nothing --minimum-gas-prices=0ujkl"

echo "waiting for node to warm up..."
sleep 20

echo "starting providers..."

screen -d -m -L -Logfile "provider0.log" -S "provider0" bash -c "./scripts/debug/restart-provider.sh 54f86a701648e8324e920f9592c21cc591b244ae46eac935d45fe962bba1102c jkl1xclg3utp4yuvaxa54r39xzrudc988s82ykve3f 0"
screen -d -m -L -Logfile "provider1.log" -S "provider1" bash -c "./scripts/debug/restart-provider.sh a29c5f0033606d1ac47db6a3327bc13a6b0c426dbfe5c15b2fcd7334b4165033 jkl1tcveayn80pe3d5wallj9kev3rfefctsmrqf6ks 1"
screen -d -m -L -Logfile "provider2.log" -S "provider2" bash -c "./scripts/debug/restart-provider.sh a490cb438024cddca16470771fb9a21938c4cf61176a46005c6a7b25ee25a649 jkl1eg3gm3e3k4dypvvme26ejmajnyvtgwwlaaeu2y 2"
screen -d -m -L -Logfile "provider3.log" -S "provider3" bash -c "./scripts/debug/restart-provider.sh 6c8a948c347079706e404ab48afc5f03203556e34ea921f3b132f2b2e9bcc87d jkl1ga0348r8zhn8k4xy3fagwvkwzvyh5lynxr5kak 3"
screen -d -m -L -Logfile "provider4.log" -S "provider4" bash -c "./scripts/debug/restart-provider.sh 8144389a23c6535e276068ff9043b2b6ff95aa3c103c35486c8f2d2363606fd5 jkl18encuf0esmxv3pxqjqvn0u4tgd6yzuc8urzlp0 4"
screen -d -m -L -Logfile "provider5.log" -S "provider5" bash -c "./scripts/debug/restart-provider.sh 0e019088a0fafa8f77cb5c0d0f6cb6b63a0015f20d2450480cbcdee44d170aab jkl1sqt9v0zwwx362szrek7pr3lpq29aygw06hgyza 5"
screen -d -m -L -Logfile "provider6.log" -S "provider6" bash -c "./scripts/debug/restart-provider.sh adf5a86ac54146b172c20b865c548e900c51439c3723af14aeab668ccd2b8ecf jkl1yu099xns2qpslvyrymxq3hwrqhevs7qxksvu8p 6"
#bash -c "./scripts/debug/debug-provider.sh adf5a86ac54146b172c20b865c548e900c51439c3723af14aeab668ccd2b8ecf jkl1yu099xns2qpslvyrymxq3hwrqhevs7qxksvu8p 6" 

read -rsp $'Press any key to end...\n' -n1 key

killall screen canined jprovd
