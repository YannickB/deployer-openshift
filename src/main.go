
package main

import (
	"os"
	"bytes"
	"os/exec"
	// "bytes"
	"fmt"
	// "net/http"
	"encoding/json"
	// "io/ioutil"
	// "crypto/tls"
	// "log"
	// "encoding/json"
	"strings"
)

func execute(name string, args ...string) {
	fmt.Printf("\n%v %v\n", name, strings.Join(args, " "))
	cmd := exec.Command(name, args...)
	// var out bytes.Buffer
	cmd.Stdout = os.Stdout
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", stdErr.String())
		fmt.Printf("%v\n", err)
		// log.Fatal(err)
	}
	// fmt.Printf("%v\n", out.String())
}

func serviceDeploy() {
	template := os.Getenv("TEMPLATE") // "node-service-template"
	application := os.Getenv("APPLICATION") // "wikicompare-api"
	sourceRepository := os.Getenv("SOURCE_REPOSITORY") // "git@gitlab.com:yannick.buron/wikicompare_api.git"
	sourceSecret := os.Getenv("SOURCE_SECRET") // "gitlab"
	sourceImage := os.Getenv("SOURCE_IMAGE") // "node-builder"

	options := ""
	for _, e := range os.Environ() {
		s := strings.Split(e, "=")
		k := s[0]
		v := s[1]
		if strings.HasPrefix(k, "OPTION_") {
			options = options + fmt.Sprintf(" %s=\"%s\"", strings.Replace(k, "OPTION_", "", -1), v)
		}
	}
	execute("sh", "-c", fmt.Sprintf("oc process %s APPLICATION_NAME='%s' SOURCE_REPOSITORY='%s' SOURCE_IMAGE='%s' SOURCE_SECRET='%s' %s | oc create -f -", template, application, sourceRepository, sourceImage, sourceSecret, options))

}

func servicePurge() {
	application := os.Getenv("APPLICATION") // "wikicompare-api"

	execute("oc", "delete", "route", application)
	execute("oc", "delete", "service", application)
	execute("oc", "delete", "deploymentconfig", application)
}

func main() {

	// host := os.Getenv("OC_HOST") // "console.infra2.wikicompare.org:8443"
	// user := os.Getenv("OC_USER") // "admin"
	// password := os.Getenv("OC_PASSWORD") // "admin"

	execute("oc", 
		"login", os.Getenv("OC_HOST"), 
		"-u", os.Getenv("OC_USER"),
		"-p", os.Getenv("OC_PASSWORD"),
		"--insecure-skip-tls-verify=True")

	switch action := os.Getenv("ACTION"); action {
	case "serviceDeploy":
		serviceDeploy()
	case "servicePurge":
		servicePurge()
	}

	msg := map[string]string{"msg": ("Hello, test!")}
	res, _ := json.Marshal(msg)
	fmt.Println(string(res))
  

	// tr := &http.Transport{
    //     TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr}

	// template, err := ioutil.ReadFile("template.json")
	// if err != nil {
	// 	fmt.Printf("Could not open file %s\n", err)
	// }

	// // jsonData := map[string]string{"name": "test"}
    // // jsonValue, _ := json.Marshal(jsonData)
	// // req, err := http.NewRequest("POST", "https://console.infra2.wikicompare.org:8443/oapi/v1/projects", bytes.NewBuffer(jsonValue))
	// req, err := http.NewRequest("POST", "https://console.infra2.wikicompare.org:8443/oapi/v1/templates", bytes.NewBuffer(template))
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer HpCE5pdLMcMaVfp2HRXU4W0w1qAAUwlBPI3o3i9VV04")

	// resp, err := client.Do(req)
    // if err != nil {
    //     fmt.Printf("The HTTP request failed with error %s\n", err)
    // } else {
    //     data, _ := ioutil.ReadAll(resp.Body)
    //     fmt.Println(string(data))
    // }
}
