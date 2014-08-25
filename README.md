goutils
=======

a repo for utlitlty functions

implements:
string functions:
	SetPrefix(s string, prefix string) string{}
	SetSuffix(s string, suffix string) string{}

path functions:
	PathExists(p string) (bool, error) {}
	Dir.Walk(path string) error {}
		Recursively walks the path, adding files that it encounters to
		an array of fileinfo. This results in creating a list of all 
		children of the path it receives.

Copyright 2014, all rights reserved.
Licensed under The Mit License. Please check the included LICENSE file for more details.	
