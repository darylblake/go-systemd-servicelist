package go_systemd_servicelist

import (
	"log"
	"os/exec"
	"strings"
)

type ServiceItems struct {

}

func CollectServiceInfo() error {

	cmd := exec.Command("systemctl", "list-units", "--type=service")
	outputData, err := cmd.Output()
	if err != nil {
		return err
	}

	log.Println(string(outputData))

	serviceItemsList := make([]ServiceItems, 0)
	err = processOutputBytesteam( outputData, &serviceItemsList )
	if err != nil {
		return err
	}

	return nil
}


func processOutputBytesteam(bytestream []byte, d *[]ServiceItems) error {
	bsString := string(bytestream)

	lines := strings.Split(bsString, "\n")
	log.Println(bsString)


	for k, v := range lines {
		if k == 0 { //header ?? skip
			log.Println("Skipping Header Row :P")
			log.Println(v)
		}

		//todo rest of data...

		//regex..?
	}

	return nil
}