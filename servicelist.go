package go_systemd_servicelist

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type ServiceItems struct {
	Name 				string			`serviceName`
	Loaded   			string			``
	State 				string
	Status  			string
	Description 		string
}

func CollectServiceInfo() ([]ServiceItems, error) {
	serviceItemsList := make([]ServiceItems, 0)
	cmd := exec.Command("systemctl", "list-units", "--type=service")
	outputData, err := cmd.Output()
	if err != nil {
		return serviceItemsList, err
	}

	log.Println(string(outputData))

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
			log.Println("Skipping Header Row :P")
			continue
		}

		//Process Service List
		items := strings.Split(v," ")
		itemLength := len(items)

		if itemLength > 4 {
			si := ServiceItems {
				Name: items[0],
				Loaded: items[1],
				State: items[2],
				Status: items[3],
			}
			desc := ""
			for i:=4; i < itemLength; i++ {

				desc = fmt.Sprintf("%s %s",desc,items[i])
			}
			si.Description = strings.Trim(desc," \n")
			*serviceItemsList = append(*serviceItemsList, si)
		} else { //We have ran off the end of the service list. this should be an empty line.
			return nil
		}
	}

	return nil
}