package services

import (
	"bytes"
	"fmt"
	"os/exec"
)

// update run the update at liquibase
func update(hash string, repository string, loglevel string) {

	repoName := repository + "_" + hash + ".log"

	cmd := exec.Command("/home/runner/work/liquibase/liquibase", "--output-file="+repoName, "--loglevel="+loglevel, "update")

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	if err := cmd.Run(); err != nil {

	} else {

		output := cmdOutput.Bytes()
		message, _ := getOutput(output)

		fmt.Println(message)

	}
}

// tag excutes the command tag of liquibase
func tag(hash string) {

	cmd := exec.Command("/home/runner/work/liquibase/liquibase", "tag", hash)

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	if err := cmd.Run(); err != nil {

	} else {

		output := cmdOutput.Bytes()
		message, _ := getOutput(output)

		fmt.Println(message)

	}

}
