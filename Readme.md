# Fedora / Redhat Systemd Service List processor.

This is a module that can extract the running services and their status' from systemd into data structures.

Tested with RHEL / Fedora 

//todo Test Amazon EC2 Linux

### To use, add as a module:

```
$ go get -u github.com/darylblake/go-systemd-servicelist@v0.0.3
```

Sample Application:

```
package main

import (
        "encoding/json"
        "fmt"
	servicelist "github.com/darylblake/go-systemd-servicelist"
)

func main() {
	
        data, _  := servicelist.CollectServiceInfo()
        
       	output, err := json.Marshal(data)
        if err != nil {
                fmt.Println("error marshalling to json")
        }
        
       	fmt.Println(string(output))
}
```

Output

```
[
{"serviceName":" abrt-journal-core.service","loaded":"loaded","state":"active","status":"running","description":"Creates ABRT problems from coredumpctl messages"},
{"serviceName":" abrt-oops.service","loaded":"loaded","state":"active","status":"running","description":"ABRT kernel log watcher"},
{"serviceName":" abrt-xorg.service","loaded":"loaded","state":"active","status":"running","description":"ABRT Xorg log watcher"},
{"serviceName":" abrtd.service","loaded":"loaded","state":"active","status":"running","description":"ABRT Automated Bug Reporting Tool"},
{"serviceName":" atd.service","loaded":"loaded","state":"active","status":"running","description":"Deferred execution scheduler"}
...
]
```


Hopefully someone finds it useful. I have....


