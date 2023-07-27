package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const urll = "https://lco.dev"
const myurl = "https:lco.dev:3000/learn?coursename=reactjs&paymentid=fjfgeiso"

func main() {
	fmt.Println("Web Requests")
	response, err := http.Get(urll)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type :%T\n", response)
	// it is important to close all the requests
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

	// parsing the url to get all the details
	result, _ := url.Parse(myurl)

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// these are key value pairs and we can extract values based like maps
	qparams := result.Query()
	fmt.Printf("type of queries is %T", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params is ", val)
	}

	// to make the url from parts, reverse of above request

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

	//get request
	// performGetRequest()

	// post request
	// performPostRequest()

	// postform request
	// performPostFromReq()

	// json important things
	coursesDone := []courses{
		{"Reactjs", 199, "tinki", "180 hrs", []string{"frontend", "starting"}},
		{"MERN", 299, "tinkiy", "182 hrs", []string{"full stack", "starting"}},
		{"Go", 99, "tinkiey", "18 hrs", []string{"frontend", "starting"}},
	}

	// it is not possible or difficiult to read this json
	// finalJson,_ := json.Marshal(coursesDone);

	// to overcome above problem use this
	//this function indent our json with spaces as third param
	finalJson, _ := json.MarshalIndent(coursesDone, "", "\t")

	fmt.Printf("%s", finalJson)
	decodeJson();
}

type courses struct {
	Name  string `json:"coursename"`
	Price int
	// this means we call the platform as website in result
	platform string `json:"website"`
	// this means this value should not be reflected on the result
	time string `json:"-"`
	// this means all the json with value as nil will be omitted in the result
	tags []string `json:"tags,omitempty"`
}

func performGetRequest() {
	const geturl = "https://jsonplaceholder.typicode.com/posts/1"

	response, err := http.Get(geturl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Cocde: ", response.StatusCode)
	fmt.Println(response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	// another way to change bytes into strings is by stringbuilder
	var responseString strings.Builder
	byteCnt, _ := responseString.Write(content)

	fmt.Println("ByteCnt is :", byteCnt)
	fmt.Println(responseString.String())
}

func performPostRequest() {
	const myurll = "http://localhost:8000/post"

	// fake json payload
	reqBody := strings.NewReader(`
		{
			"coursename":"reactjs",
			"price":0,
		}
	`)
	response, err := http.Post(myurll, "application/json", reqBody)

	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(content))
}

func performPostFromReq() {
	const myurll = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "aman")
	data.Add("lastname", "manwani")
	data.Add("email", "ahfgf")

	response, _ := http.PostForm(myurll, data)

	content, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(content))
}

func decodeJson() {
	jsonData := []byte(`
	{
		"coursename": "Reactjs",
		"Price": 199
	}
	`)

	var coursess courses

	checkValid := json.Valid(jsonData);
	if checkValid {
		fmt.Println("json was valid");
		json.Unmarshal(jsonData,&coursess)
		fmt.Printf("%#v\n",coursess);
	}else{
		fmt.Println("json was not vallid");
	}

	// some cases where you just want to add data in form of key value pairs
	var keyValJson map[string]interface{}
	json.Unmarshal(jsonData,&keyValJson);
	fmt.Printf("%#v\n",keyValJson);

	//  to get all the key value pairs , we can iterate through loops
	for k,v := range keyValJson {
		fmt.Printf("key is %v and value is %v",k,v);
	}
}
