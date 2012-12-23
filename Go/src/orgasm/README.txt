Checkout the following packages with "go get"

github.com/ziutek/mymysql/godrv
github.com/ziutek/mymysql/thrsafe
github.com/astaxie/beedb

When those packages are installed, build the following package from the PU source tree
"go install github/astaxie/beedb"
Copy the output file (beedb.a) from the local pkg directory (eg. pokemon-universe\Go\pkg\windows_386\github\astaxie)
  to your global go pkg dir (eg. C:\Go\pkg\windows_386\github.com\astaxie)

Build the "ORGASM" project like you always do. This should automatically build the "nonamelib" and "pulogic" packages.

If everything went correctly you should be albe to run the server.