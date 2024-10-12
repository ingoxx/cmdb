package handle

import (
	"github.com/Lxb921006/cmdb/project/controllers/web"
	"os/exec"
)

type Handle struct {
	Item string
}

func NewHandle(item string) (result string, err error) {
	h := Handle{Item: item}
	return h.RefreshCdn()
}

func (h *Handle) executeCmd(cmd string) (result string, err error) {
	out, err := exec.Command("sh", cmd, h.Item).Output()
	if err != nil {
		return string(out), err
	}

	return
}

func (h *Handle) RefreshCdn() (result string, err error) {
	result, err = h.executeCmd(web.RefreshAwsCdnScript)
	if err != nil {
		return result, err
	}

	return result, nil
}
