
Hello !
It will be  client server-application which provides following features:

From client side:
 * watching for changes in specified txt-files. in my case it will repository configuration files as example:
   (/etc/apt/sources.d/* ; /etc/zypp.d/* ;  )
   Files that include information about installed packages will also being monitored . As example ( /var/lib/dpkg/status or /var/lib/rpm/* )
  * Sendind information about configured repositories and installed packages to server .
From server side:
  * Retrieving from clients information about repositories and packages
  * Monitoring modification time of each  Packages.gz or primary.xml.gz . This packages usualy store information about available packages in remote repository  
