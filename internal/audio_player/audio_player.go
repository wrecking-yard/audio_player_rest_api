package audio_player

import (
	"os/exec"
)

type Result struct {
	Log []byte
	Err error
}

type Remote struct {
	Cmd  string
	Args []string
	Result Result
}

func (r *Remote) Run() {
	cmd := exec.Command(r.Cmd, r.Args...)
	output, err := cmd.CombinedOutput()
	r.Result.Err = err
	r.Result.Log = output
}

func NewRemote(cmd string, args []string) Remote {
	_cmd := "/usr/bin/rhythmbox-client"
	if cmd != "" {
		_cmd = cmd
	}
	remote := Remote{
		Cmd: _cmd,
		Args: args,
	}

	return remote
}
