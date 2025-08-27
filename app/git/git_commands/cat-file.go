package gitcommands

//git cat-file -p hash
func gitcatfile(flag string, commitSHA string) {
	
	switch flag {
	case "-p" :
		folder := commitSHA[:2]
		file := commitSHA[2:]

		completePath := 
	}
}
