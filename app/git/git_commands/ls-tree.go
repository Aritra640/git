package gitcommands

import "os"

type Git_ls_tree struct {
	flag string
	sha  string
}

func (g *Git_ls_tree) Init() {
	if len(os.Args) == 4 {
		g.flag = os.Args[2]
		g.sha = os.Args[3]

	} else {
		g.flag = ""
		g.sha = os.Args[2]
	}
}


func (g *Git_ls_tree) Execute() {

	folder := g.sha[:2]
	file := g.sha[2:]


}
