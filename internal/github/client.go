package github

import (
	"fmt"
	"io"
	"net/http"
)

func ProcessDiff(diffFileUrl string) {
	resp, err := http.Get(diffFileUrl)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	diffByte, err := io.ReadAll(resp.Body)

	diff := string(diffByte)

	fmt.Println(diff)
}
