# Installing the Dependencies #

First of all, PU is programmed in [the Go Programming Language](http://golang.org/). Follow [Getting Started](http://golang.org/doc/install.html) in order to install Go on your machine.

We also use the newest SDL, SDL\_ttf and SDL\_image. These require other libraries too, let's install them:
```
sudo apt-get install libgl1-mesa-dev libpng12-dev libjpeg62-dev libfreetype6-dev autoconf subversion

hg clone http://hg.libsdl.org/SDL -r 1281a3f1f0a6
cd SDL
./autogen.sh
./configure
make
sudo make install

hg clone http://hg.libsdl.org/SDL_ttf -r a6e04cc57348
cd SDL_ttf
./autogen.sh
./configure
make
sudo make install

hg clone http://hg.libsdl.org/SDL_image -r b1c1ec3a8d49
cd SDL_image
./autogen.sh
./configure
make
sudo make install
```


# Compiling! #

Let's checkout the newest PU:
```
svn checkout https://pokemon-universe.googlecode.com/svn/trunk/ pu
```

In order for the client and server to compile, you must install the packages in the `Packages` directory. However, calling `make` will automate this together with compiling the client and server.
```
cd pu
make
```