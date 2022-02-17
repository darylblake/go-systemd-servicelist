package go_systemd_servicelist

import (
	"log"
	"os/exec"
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
		log.Println(items)

		/*
		if itemLength > 4 {

			for char = 0; char < len(items)

			si := ServiceItems {
				Name: items[2],
				Loaded: items[1],
				State: ,
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

		 */
	}

	return nil
}