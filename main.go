
package main

import (
	"os"
	"bytes"
	"os/exec"
	// "bytes"
	"fmt"
	// "net/http"
	// "encoding/json"
	// "io/ioutil"
	// "crypto/tls"
	// "log"
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

func main() {

	host := "console.infra2.wikicompare.org:8443"
	user := "admin"
	password := "admin"
	template := "node-service-template"
	application := "wikicompare-api"
	database := "postgres"
	adminPassword := "MYPASSWORD"
	sourceRepositoryUrl := "git@gitlab.com:yannick.buron/wikicompare_api.git"
	sourceImage := "node-builder"
	sourceSecret := "gitlab"


	execute("oc", "login", host, "-u", user, "-p", password, "--insecure-skip-tls-verify=True")
	execute("sh", "-c", fmt.Sprintf("oc process %s APPLICATION_NAME='%s' POSTGRESQL_HOST='%s' ADMIN_PASSWORD='%s' SOURCE_REPOSITORY_URL='%s' SOURCE_IMAGE='%s' SOURCE_SECRET='%s' | oc create -f -", template, application, database, adminPassword, sourceRepositoryUrl, sourceImage, sourceSecret))

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
