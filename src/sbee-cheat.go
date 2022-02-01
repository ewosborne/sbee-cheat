// port to go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {
	var words map[string][]string

	resp, err := http.Get("https://www.nytimes.com/puzzles/spelling-bee")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	today, err := regexp.Compile(`"today":({.*?})`)
	if err != nil {
		log.Fatal(err)
	}

	to_json := today.FindSubmatch(body)[1]

	json.Unmarshal(to_json, &words)

	fmt.Printf("pangrams %v\n", words["pangrams"])
	fmt.Printf("answers %v (%d)\n", words["answers"], len(words["answers"]))

	// 	res = {
	// 	'pangrams': [x for x in re.findall('\w+', m.group(1))],
	// 	'answers': [x for x in re.findall('\w+', m.group(2))]
	// }

}
