package main

import (
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
	response,err := http.Get(urll)
	
	if err!=nil {
		panic(err);
	}

	fmt.Printf("Response is of type :%T\n",response);
	// it is important to close all the requests
	defer response.Body.Close()

	databytes,err := ioutil.ReadAll(response.Body);
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content);

	// parsing the url to get all the details 
	result,_ := url.Parse(myurl)

	fmt.Println(result);
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery);

	// these are key value pairs and we can extract values based like maps
	qparams := result.Query();
	fmt.Printf("type of queries is %T",qparams)

	fmt.Println(qparams["coursename"]);

	for _,val := range qparams {
		fmt.Println("Params is ", val);
	}

	// to make the url from parts, reverse of above request

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	} 

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL);

	//get request
	performGetRequest();

}

func performGetRequest(){
	const geturl = "https://jsonplaceholder.typicode.com/posts/1"

	response,err := http.Get(geturl);

	if err!=nil {
		panic(err);
	}
	defer response.Body.Close();

	fmt.Println("Status Cocde: ", response.StatusCode);
	fmt.Println(response.ContentLength);

	content,_ := ioutil.ReadAll(response.Body);
	fmt.Println(string(content));

	// another way to change bytes into strings is by stringbuilder
	var responseString strings.Builder
	byteCnt,_ := responseString.Write(content);

	fmt.Println("ByteCnt is :", byteCnt);
	fmt.Println(responseString.String())
}
