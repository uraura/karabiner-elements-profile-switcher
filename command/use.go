package command

import (
	"github.com/urfave/cli"
	"os"
	"io/ioutil"
	"github.com/docker/machine/libmachine/log"
	"encoding/json"
)

func CmdUse(c *cli.Context) {
	jsonPath := os.Getenv("HOME") + "/.config/karabiner/karabiner.json"

	in, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	var karabinerConfig map[string]*interface{}
	json.Unmarshal(in, &karabinerConfig)

	karabinerProfiles := (*karabinerConfig["profiles"]).([]interface{})
	for _, profile := range karabinerProfiles {
		m := profile.(map[string]interface{})
		if m["name"] == c.Args().First() {
			m["selected"] = true
		} else {
			m["selected"] = false
		}
	}

	out, _ := json.Marshal(karabinerConfig)

	ioutil.WriteFile(jsonPath, out, 0644)
}
