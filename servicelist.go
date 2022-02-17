package go_systemd_servicelist

import (
	"os/exec"
	"regexp"
	"strings"
)

type ServiceItems struct {
	Name 				string			`json:"serviceName"`
	Loaded   			string			`json:"loaded"`
	State 				string			`json:"state"`
	Status  			string			`json:"status"`
	Description 		string			`json:"description"`
}

func CollectServiceInfo() ([]ServiceItems, error) {
	serviceItemsList := make([]ServiceItems, 0)
	cmd := exec.Command("systemctl", "list-units", "--type=service")
	outputData, err := cmd.Output()
	if err != nil {
		return serviceItemsList, err
	}


	err = processOutputBytesteam(outputData, &serviceItemsList)
	if err != nil {
		return serviceItemsList, err
	}

	return serviceItemsList, nil
}


func processOutputBytesteam(bytestream []byte, serviceItemsList *[]ServiceItems) error {

	bsString := string(bytestream)

	lines := strings.Split(bsString, "\n")

	for k, v := range lines {
		if k == 0 { //header ?? skip
			continue
		}

		re := regexp.MustCompile(`\s(.+\.service)\s+([A-z]+)\s+([A-z]+)\s+([A-z]+)\s+(.+)`)
		segments := re.FindAllStringSubmatch(v, -1)

		if len(segments) > 0 {
			si := ServiceItems {
				Name: segments[0][1],
				Loaded: segments[0][2],
				State: segments[0][3],
				Status: segments[0][4],
				Description: segments[0][5],
			}
			*serviceItemsList = append(*serviceItemsList, si)
		}
	}

	return nil
}