package gh

import (
	"os/exec"
)

func GetRepos() {
	ghRepoListCommand := "gh repo list | fzf"
	cmd := exec.Command("tmux", "bind-key", "T", "new-window", "-n", "repos", ghRepoListCommand)
	cmd.Start()
}
