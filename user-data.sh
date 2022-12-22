Content-Type: multipart/mixed; boundary="//"
MIME-Version: 1.0

--//
Content-Type: text/cloud-config; charset="us-ascii"
MIME-Version: 1.0
Content-Transfer-Encoding: 7bit
Content-Disposition: attachment; filename="cloud-config.txt"

#cloud-config
cloud_final_modules:
- [scripts-user, always]

--//
Content-Type: text/x-shellscript; charset="us-ascii"
MIME-Version: 1.0
Content-Transfer-Encoding: 7bit
Content-Disposition: attachment; filename="userdata.txt"

#!/bin/bash
# CUSTOM: Set your URL to download the built rapidchain binary
wget -O /home/ubuntu/rapidchain https://xshard-rapidchain.s3.amazonaws.com/rapidchain
chmod a+x /home/ubuntu/rapidchain
ulimit -n 65535
sysctl -w net.ipv4.tcp_tw_reuse=1
# CUSTOM: Options
/home/ubuntu/rapidchain
--//--
