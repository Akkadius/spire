package boot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type EQEmuConfigJson struct {
	Server struct {
		Zones struct {
			Defaultstatus string `json:"defaultstatus"`
			Ports         struct {
				Low  string `json:"low"`
				High string `json:"high"`
			} `json:"ports"`
		} `json:"zones"`
		Qsdatabase struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
			Db       string `json:"db"`
		} `json:"qsdatabase"`
		Chatserver struct {
			Port string `json:"port"`
			Host string `json:"host"`
		} `json:"chatserver"`
		Mailserver struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"mailserver"`
		World struct {
			Loginserver1 struct {
				Account  string `json:"account"`
				Password string `json:"password"`
				Legacy   string `json:"legacy"`
				Host     string `json:"host"`
				Port     string `json:"port"`
			} `json:"loginserver1"`
			Loginserver2 struct {
				Port     string `json:"port"`
				Account  string `json:"account"`
				Password string `json:"password"`
				Host     string `json:"host"`
			} `json:"loginserver2"`
			TCP struct {
				IP   string `json:"ip"`
				Port string `json:"port"`
			} `json:"tcp"`
			Telnet struct {
				IP      string `json:"ip"`
				Port    string `json:"port"`
				Enabled string `json:"enabled"`
			} `json:"telnet"`
			Key       string `json:"key"`
			Shortname string `json:"shortname"`
			Longname  string `json:"longname"`
		} `json:"world"`
		Database struct {
			Db       string `json:"db"`
			Host     string `json:"host"`
			Port     string `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"database"`
		Files struct {
			Opcodes     string `json:"opcodes"`
			MailOpcodes string `json:"mail_opcodes"`
		} `json:"files"`
		Directories struct {
			Patches string `json:"patches"`
			Opcodes string `json:"opcodes"`
		} `json:"directories"`
	} `json:"server"`
}

const eqemuConfigJson = "eqemu_config.json"

func getEQEmuConfig() EQEmuConfigJson {
	if _, err := os.Stat(eqemuConfigJson); err == nil {
		// fmt.Printf("Reading from config [%v]\n", eqemuConfigJson)
		body, err := ioutil.ReadFile(eqemuConfigJson)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}

		config := EQEmuConfigJson{}
		err = json.Unmarshal(body, &config)
		if err != nil {
			log.Fatalf("unable to unmarshal file [%v] error [%v]", eqemuConfigJson, err)
		}

		return config
	}

	return EQEmuConfigJson{}
}
