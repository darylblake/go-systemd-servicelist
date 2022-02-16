package go_systemd_servicelist

import (
	"log"
	"os/exec"
)

func CollectServiceInfo() {

	cmd := exec.Command("systemctl", "list-units", "--type=service")
	outputData, err := cmd.Output()
	if err != nil {
		return
	}

	log.Println(outputData)

}
