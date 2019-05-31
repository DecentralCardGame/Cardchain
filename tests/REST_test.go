package cardservice

import(
	"fmt"
	"bytes"
	"os/exec"
	"net/http"
	//"net/http/httptest"
	"testing"
	"io/ioutil"
	"encoding/json"
	"github.com/xeipuuv/gojsonschema"
)

type TestCases struct {
		Prevaluate []string `json:"prevaluate"`
    TestCases []TestCase `json:"testCases"`
		Postvaluate []string `json:"postvaluate"`
}

type TestCase struct {
    Name   		string	`json:"name"`
    Route   	string	`json:"route"`
    ReqBody		interface{}	`json:"reqBody"`
		Evaluate	string 	`json:"evaluate"`
		Method		string  `json:"method"`
    Expected 	interface{}	`json:"expected"`
}


func TestRESTroutes(t *testing.T) {

	// read the file with the test specs
	file, _ := ioutil.ReadFile("REST_test.json")


	data := TestCases{}

	_ = json.Unmarshal([]byte(file), &data)



	cmd := exec.Command("sleep", "1")
	fmt.Println("Running command and waiting for it to finish...")
	err := cmd.Run()
	fmt.Println("Command finished with error: ", err)

	/*
	for i := 0; i < len(data.Prevaluate); i++ {
		fmt.Println(data.Prevaluate[i])
	}

	for i := 0; i < len(data.Postvaluate); i++ {
		fmt.Println(data.Postvaluate[i])
	}


	for i := 0; i < len(data.TestCases); i++ {
		fmt.Println("Name: ", data.TestCases[i].Name)
		fmt.Println("Route: ", data.TestCases[i].Route)
		fmt.Println("ReqBody: ", data.TestCases[i].ReqBody)
		fmt.Println("Evaluate: ", data.TestCases[i].Evaluate)
		fmt.Println("Expected: ", data.TestCases[i].Expected)
	}
	*/
	i := 1

	fmt.Println("before parse: ", data.TestCases[i].ReqBody)
	jsonString, _ := json.Marshal(data.TestCases[i].ReqBody)
	fmt.Println("parsed back: ", string(jsonString))

	message := fmt.Sprintf(string(jsonString), "cosmos1305dedj96kkkkhmruf5yzj82w63d5y73p7wr8l","cosmos1305dedj96kkkkhmruf5yzj82w63d5y73p7wr8l")
	fmt.Println("parsed on: ", string(message))

	t.Run(data.TestCases[i].Name, func(t *testing.T) {

		var body []byte

		switch data.TestCases[i].Method {

    case "GET":
			res, err := http.Get(data.TestCases[i].Route)
			if err != nil {
				t.Errorf(fmt.Sprintf("GET request failed"))
			}
			defer res.Body.Close()
			body, err = ioutil.ReadAll(res.Body)

    case "POST":



			fmt.Println("http POST: ", data.TestCases[i].Route, message)

			res, err := http.Post(data.TestCases[i].Route, "application/json", bytes.NewBuffer([]byte(message)))
			if err != nil {
				t.Errorf(fmt.Sprintf("POST request failed"))
			}




			defer res.Body.Close()
			body, err = ioutil.ReadAll(res.Body)
    case "PUT":
      fmt.Println("three")
		default:
	    t.Errorf(fmt.Sprintf("Test method unknown - %s", data.TestCases[i].Method))
	  }

		//body = []byte("bla")

		fmt.Println(string(body))

		// Check the response body is what we expect.
		schemaLoader := gojsonschema.NewGoLoader(data.TestCases[i].Expected)
		documentLoader := gojsonschema.NewStringLoader(string(body))

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
				panic(err.Error())
		}

		if result.Valid() {
				fmt.Printf("The document is valid\n")
		} else {
				fmt.Printf("The document is not valid. see errors :\n")
				for _, desc := range result.Errors() {
						//fmt.Printf("- %s\n", desc)
						t.Errorf(fmt.Sprintf("%s", desc))
				}
		}

	})

}
