package cardservice

import(
	"fmt"
	"log"
	"time"
	"bytes"
	"strings"
	//"os"
	"os/exec"
	//"os/signal"
	"net/http"
	//"net/http/httptest"
	"testing"
	"io/ioutil"
	"encoding/json"
	"github.com/xeipuuv/gojsonschema"
)

type TestCases struct {
		Prevaluate []string `json:"prevaluate"`
		RunProcesses []string `json:"runProcesses"`
    TestCases []TestCase `json:"testCases"`
		Postvaluate []string `json:"postvaluate"`
}

type TestCase struct {
    Name   		string	`json:"name"`
    Route   	string	`json:"route"`
    ReqBody		interface{}	`json:"reqBody"`
		Evaluate	[]string 	`json:"evaluate"`
		Method		string  `json:"method"`
		Sign			string  `json:"sign"`
		Broadcast	string  `json:"broadcast"`
    Expected 	interface{}	`json:"expected"`
}

/*
func getFireSignalsChannel() chan os.Signal {

  c := make(chan os.Signal, 1)
  signal.Notify(c,
    // https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
    syscall.SIGTERM, // "the normal way to politely ask a program to terminate"
    syscall.SIGINT, // Ctrl+C
    syscall.SIGQUIT, // Ctrl-\
    syscall.SIGKILL, // "always fatal", "SIGKILL and SIGSTOP may not be caught by a program"
    syscall.SIGHUP, // "terminal is disconnected"
  )
  return c
}

func exit() {
  syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
}
*/

func TestRESTroutes(t *testing.T) {

	/*
	exitChan := getFireSignalsChannel()
  input, err := os.Open("input.txt")
  if err != nil {
    panic(err)
  }
  defer input.Close()
  <-exitChan
  fmt.Println("Exiting!")
	fmt.Println("Exiting!")
	fmt.Println("Exiting!")
  return
	*/

	// read the file with the test specs
	file, err := ioutil.ReadFile("REST_test.json")
	if err != nil {
		t.Errorf(fmt.Sprintf("reading REST_test.json failed"))
	}


	data := TestCases{}

	_ = json.Unmarshal([]byte(file), &data)

	// run prevaluate commands
	for i := 0; i < len(data.Prevaluate); i++ {
		evaluateCmd(data.Prevaluate[i])
	}

	background := processes{}

	// run prevaluate commands
	for i := 0; i < len(data.RunProcesses); i++ {
		background.startProcess(data.RunProcesses[i])
	}
	defer background.terminate()

	// run all of the test cases
	for i := 0; i < len(data.TestCases); i++ {

		t.Run(data.TestCases[i].Name, func(t *testing.T) {
			thisCase := data.TestCases[i]

			log.Println("\x1b[33;1mrunning test: \x1b[0m", thisCase.Name)

			// evaluate test specific commands
			var evals []interface{}
			for i := 0; i < len(thisCase.Evaluate); i++ {
				resultval := evaluateCmd(thisCase.Evaluate[i])
				resultval = strings.Replace(resultval,"\n","",-1)
				evals = append(evals, resultval)
			}


			var resBody []byte

			switch thisCase.Method {
	    case "GET":
				// insert values into request address
				thisCase.Route = fmt.Sprintf(thisCase.Route, evals...)

				log.Println("\x1b[36;1mHTTP-GET: \x1b[0m", thisCase.Route)

				res, err := http.Get(thisCase.Route)
				if err != nil {
					t.Errorf(fmt.Sprintf("\x1b[31;1m GET request failed \x1b[0m"))
				}
				defer res.Body.Close()
				resBody, err = ioutil.ReadAll(res.Body)

	    case "POST":
				// insert values into request body
				reqbody, err := json.Marshal(thisCase.ReqBody)
				if err != nil {
					t.Errorf(fmt.Sprintf("\x1b[31;1m thisCase.ReqBody marshal failed \x1b[0m"))
				}
				reqbodyStr := fmt.Sprintf(string(reqbody), evals...)

				log.Println("\x1b[36;1mHTTP-POST: \x1b[0m", thisCase.Route, reqbodyStr)

				res, err := http.Post(thisCase.Route, "application/json", bytes.NewBuffer([]byte(reqbodyStr)))
				if err != nil {
					t.Errorf(fmt.Sprintf("\x1b[31;1m POST request failed \x1b[0m"))
				}

				defer res.Body.Close()
				resBody, err = ioutil.ReadAll(res.Body)

	    case "PUT":
				// insert values into request body
				reqbody, err := json.Marshal(thisCase.ReqBody)
				if err != nil {
					t.Errorf(fmt.Sprintf("\x1b[31;1m thisCase.ReqBody marshal failed \x1b[0m"))
				}
				reqbodyStr := fmt.Sprintf(string(reqbody), evals...)

				log.Println("\x1b[36;1mHTTP-PUT: \x1b[0m", thisCase.Route, reqbodyStr)

				client := &http.Client{}

				req, err := http.NewRequest("PUT", thisCase.Route, bytes.NewBuffer([]byte(reqbodyStr)))
				if err != nil {
					t.Errorf(fmt.Sprintf("\x1b[31;1m PUT request construction failed \x1b[0m"))
				}
				// ...
				req.Header.Add("If-None-Match", `W/"wyzzy"`)
				res, err := client.Do(req)
				if err != nil {
					t.Errorf(fmt.Sprintf("\x1b[31;1m sending PUT request failed \x1b[0m"))
				}

				defer res.Body.Close()
				resBody, err = ioutil.ReadAll(res.Body)

			default:
		    t.Errorf(fmt.Sprintf("\x1b[31;1m Test method unknown - %s \x1b[0m", thisCase.Method))
				return
		  }

			log.Println("\x1b[36;1mresponse body: \x1b[0m", string(resBody))

			// Check if the response body is what we expect.
			schemaLoader := gojsonschema.NewGoLoader(thisCase.Expected)
			documentLoader := gojsonschema.NewStringLoader(string(resBody))
			result, err := gojsonschema.Validate(schemaLoader, documentLoader)
			if err != nil {
				log.Println("\x1b[31;1mgojsonschema.Validate failed. This can lead to data structures that make following tests fail.\x1b[0m")
				
				//for _, desc := range result.Errors() {
					//	fmt.Printf("- %s\n", desc)
				//}

				t.Errorf(fmt.Sprintf("\x1b[31;1m gojsonschema.Validate failed - %s\x1b[0m", err.Error()))
			} else if result.Valid() {
					log.Println("\x1b[32;1mThe response is valid, if possible try to sign and broadcast it.\x1b[0m")

					err = ioutil.WriteFile("unsignedTx.json", resBody, 0644)
					if err != nil {
						t.Errorf(fmt.Sprintf("\x1b[31;1m writing unsignedTx.json failed \x1b[0m"))
					}

					if thisCase.Sign != "" {
						evaluateCmd(thisCase.Sign)
					}
					var broadcast string
					if thisCase.Broadcast != "" {
						broadcast = evaluateCmd(thisCase.Broadcast)

						success := strings.Contains(broadcast, `"success":true,"log"`)
						if success {
							log.Println("\x1b[32;1mTransaction successfully signed and broadcasted.","\x1b[0m")
						} else {
							t.Errorf(fmt.Sprintf("\x1b[31;1m broadcast failed \x1b[0m"))
						}
					}
			} else {
					log.Println("\x1b[31;1m The document is not valid. see errors : \x1b[0m")
					for _, desc := range result.Errors() {
							fmt.Printf("- %s\n", desc)
							t.Errorf(fmt.Sprintf("%s", desc))
					}
			}
		})
	}

	// postvaluate

}

func evaluateCmd(cmdstr string) string {

	parts := strings.Fields(cmdstr)
  head := parts[0]
  parts = parts[1:len(parts)]
	log.Println("\x1b[36;1mevaluating: \x1b[0m", head, parts)

	cmd := exec.Command(head, parts...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
	    log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	log.Println("\x1b[36;1mresult: \x1b[0m" + out.String())

	return out.String()
}


type processes struct {
	running []*exec.Cmd
}

func (pr *processes) startProcess(cmdstr string) {

	parts := strings.Fields(cmdstr)
	head := parts[0]
	parts = parts[1:len(parts)]
	log.Println("\x1b[36;1mstart process: \x1b[0m", head, parts)

	cmd := exec.Command(head, parts...)

	//cmd.Dir = entryPath

	if err := cmd.Start(); err != nil {
	    log.Printf("\x1b[31;1mFailed to start cmd: %v\x1b[0m", err)
	    return
	}

	time.Sleep(3 * time.Second)
	pr.running = append(pr.running, cmd)
}

func (pr *processes) terminate() {
	log.Printf("terminate called!")
	for i := 0; i < len(pr.running); i++ {
		log.Println("killing process ", i)
		if err := pr.running[i].Process.Kill(); err != nil {
  		log.Fatal("\x1b[31;1mfailed to kill process: \x1b[0m", err)
		}
	}
}
