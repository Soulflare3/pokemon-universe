all:
	cd Packages; gomake install
	cd Client; gomake
	cd Server; gomake
	
clean:
	cd Packages; gomake clean
	cd Client; gomake clean
	cd Server; gomake clean
