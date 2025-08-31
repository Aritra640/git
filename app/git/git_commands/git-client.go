package gitcommands

type GitClient interface {
	Init()
	Execute()
}
