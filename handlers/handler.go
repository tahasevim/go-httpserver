//Package handlers contains all endpoints's handlers.
//All endpoints based on HTTP methods and entire handlers operate coming requests according to it.
package handlers

import(
	"net/http"
	"github.com/tahasevim/responsiveweb/templates"
	"log"
	"strconv"
	"strings"
	"encoding/json"
	"time"
	"fmt"
	"math/rand"
)


//GetHandlers adds handlers to the a map and returns it.
func GetHandlers()map[string]func(http.ResponseWriter,*http.Request){
	handlerList := make(map[string]func(http.ResponseWriter,*http.Request))
	handlerList["/"] = indexHandler
	handlerList["/ip"] = ipHandler
	handlerList["/headers"] = headersHandler
	handlerList["/get"] = getHandler
	handlerList["/user-agent"] = useragentHandler
	handlerList["/uuid"] = uuidHandler
	handlerList["/post"] = postHandler
	handlerList["/delete"] = deleteHandler
	handlerList["/put"] = putHandler
	handlerList["/anything"] = anythingHandler
	handlerList["/anything/"] = anythingHandler
	handlerList["/encoding/utf8"] = utf8Handler
	handlerList["/gzip"] = gzipHandler
	handlerList["/deflate"] = deflateHandler
	handlerList["/brotli"] = brotliHandler
	handlerList["/status/"] = statusHandler
	handlerList["/response-headers"] = responseHeaderHandler
	handlerList["/redirect/"] = redirectMultiHandler
	handlerList["/redirect-to"] = redirectToHandler
	//handlerList["/relative-redirect"] = relativeRedHandler
	//handlerList["/absolute-redirect"] = absoluteRedHandler
	handlerList["/cookies"] = cookieHandler
	handlerList["/cookies/"] = cookieSetDelhandler
	handlerList["/basic-auth"] = basicAuthHandler
	handlerList["/hidden-basic-auth"] = hiddenBasicAuthHandler
	handlerList["/stream/"] = streamHandler
	handlerList["/delay/"] = delayHandler
	handlerList["/html"] = htmlHandler
	handlerList["/robots.txt"] = robotsTextHandler
	handlerList["/deny"] = denyHandler
	handlerList["/cache"] = cacheHandler
	handlerList["/cache/"] = cacheControlHandler
	handlerList["/bytes/"] = bytesHandler
	handlerList["/links/"] = linkHandler
	handlerList["/image"] = imageHandler
	handlerList["/image/png"] = pngHandler
	handlerList["/image/jpeg"] = jpegHandler
	handlerList["/image/webp"] = webpHandler
	handlerList["/image/svg"] = svgHandler
	handlerList["/forms/post"] = formsHandler
	handlerList["/xml"] = xmlHandler
	return handlerList
}

//ipHandler handles a GET request and sends a response in JSON format that contains IP address of client which made request.
func ipHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"origin")	
	w.Write(makeJSONresponse(jsonData))
}

//indexHandler handles a GET request and sends a HTML page that contains links of endpoints.
func indexHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	templates.IndexTemplate.ExecuteTemplate(w, "index", nil)
}
//headersHandler handles a GET request and sends a response in JSON format that contains header of the coming request.
func headersHandler(w http.ResponseWriter,r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"headers")	
	w.Write(makeJSONresponse(jsonData))
}
//getHandler handles a GET request and sends a response in JSON format that contains args,IP,headers,url of the coming request.
func getHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"args","headers","origin","url")
	w.Write(makeJSONresponse(jsonData))
}

//useragentHandler handles a GET request and sends a response in JSON format that contains user-agent of the coming request.
func useragentHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"user-agent")	
	w.Write(makeJSONresponse(jsonData))
}

//uuidHandler handles a GET request and sends a response in JSON format that contains uuid (Universally unique identifier).
//uuid is obtained by operating system's "uuidgen" tool.
func uuidHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"uuid")	
	w.Write(makeJSONresponse(jsonData))
}

//postHandler handles a POST request and sends a response in JSON format that contains args,data,files,form,headers,IP,url of the coming request.
func postHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

//deleteHandler handles a DELETE request and sends a response in JSON format that contains args,data,files,form,headers,IP,url of the coming request.
func deleteHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "DELETE" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

//putHandler handles a PUT request and sends a response in JSON format that contains args,data,files,form,headers,IP,url of the coming request.
func putHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "PUT" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

//patchHandler handles a PATCH request and sends a response in JSON format that contains args,data,files,form,headers,IP,url of the coming request.
func patchHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "PATCH" {
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url")
	w.Write(makeJSONresponse(jsonData))	
}

//anythingHandler can handle any type of request and sends a response in JSON format that contains args,data,files,form,headers,IP,url,method of the coming request.
func anythingHandler(w http.ResponseWriter, r *http.Request){
	jsonData := getAllJSONdata(r ,"args","data","files","form","headers","json","origin","url","method")
	w.Write(makeJSONresponse(jsonData))
}

//utf8Handler handles a GET request and sends a UTF8 encoded template that contains a lot of different UTF8 encoded characters.
func utf8Handler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	templates.Utf8Template.ExecuteTemplate(w,"utf8",nil)
}

//gzipHandler handles a GET request and sends a response in JSON format that contains gzipped,headers,method,IP of the coming request.
func gzipHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"gzipped","headers","method","origin")
	w.Write(makeJSONresponse(jsonData))
	
}

//brotliHandler handles a GET request and sends a response in JSON format that contains gzipped,headers,method,IP of the coming request.
func brotliHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"brotli","headers","method","origin")
	w.Write(makeJSONresponse(jsonData))	
}

//deflateHandler handles a GET request and sends a response in JSON format that contains gzipped,headers,method,IP of the coming request.
func deflateHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := getAllJSONdata(r,"deflated","headers","method","origin")
	w.Write(makeJSONresponse(jsonData))	
}

//statusHandler can handle any type of request and redirects the coming request to the that status's url.
func statusHandler(w http.ResponseWriter, r *http.Request){
	stat,_ := strconv.ParseInt(r.URL.Path[len("/status/"):],10,64)
	if int(stat)==0{
		w.WriteHeader(418)
		http.Redirect(w,r,"/status/418",418)
	}
	http.Redirect(w,r,"/status/"+string(stat),int(stat))
}

//responseHeaderHandler handles a GET or POST request and sends a response in JSON format.
//It prepares a JSON response from url of the coming request.
func responseHeaderHandler(w http.ResponseWriter, r *http.Request){
	if !(r.Method == "GET" || r.Method == "POST"){
		http.Error(w,"Method Not Allowed",405)
		return
	}
	jsonData := jsonMap{}
	for key,value := range r.URL.Query(){
		if len(value) == 1{
			jsonData[key] = value[0]
		}else {
			jsonData[key] = value
		}
	}
	jsonData["Content-Type"] = "application/json"
	w.Write(makeJSONresponse(jsonData))	
}

//redirectMultiHandler handles a GET request and redirects the coming request n times.
func redirectMultiHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	ntime, err := strconv.ParseInt(r.URL.Path[len("/redirect/"):],10,64)
	if int(ntime)<0{
		w.Write([]byte("Invalid n"))
		return
	}
	if int(ntime)== 0{
		http.Redirect(w,r,"/get",302)
	}
	if err != nil {
		w.Write([]byte("Invalid n"))
		return
	}
	for i:=0;i<int(ntime);i++{
		http.Redirect(w,r,"/get",302)	
	}
}

//redirectToHandler handles a GET request and redirects the coming request to the given url parameter.
func redirectToHandler(w http.ResponseWriter, r *http.Request){
	var stat int
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	url := r.URL.Query().Get("url")
	statstr, _ := strconv.ParseInt(r.URL.Query().Get("status_code"),10,64)
	if statstr == 0{
		stat = 302
	}else{
		stat = int(statstr)
	}
	http.Redirect(w,r,url,stat)
}

//cookieHandler handles a GET request and sends a response in JSON format that contains cookies.
func cookieHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	jsonData := jsonMap{}
	cookieMap := jsonMap{}
	for _,cookie := range r.Cookies(){
		cookieMap[cookie.Name] = cookie.Value
	}
	jsonData["cookies"] = cookieMap
	w.Write(makeJSONresponse(jsonData))
}

//cookieSetDelhandler handles a GET request and sends a response in JSON format that contains cookies.
//It sets or deletes cookies according to given url (/set or /delete).
func cookieSetDelhandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	if r.URL.Path == "/cookies/set"{
		jsonData := setCooki(w,r)
		w.Write(makeJSONresponse(jsonData))
		return	
	}
	if r.URL.Path == "/cookies/delete"{
		delCooki(w,r)
		http.Redirect(w,r,"/cookies",302)
		return
	}
	http.Redirect(w,r,"/cookies",302)
}

//basicAuthHandler handles a GET request and sends a response in JSON format.
//It recieves password and username from client and checks that it is valid or not.
func basicAuthHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	jsonData := jsonMap{}
	jsonData = getAllJSONdata(r,"authenticated","user")
	if len(r.URL.String())==12{
		w.Write(makeJSONresponse(jsonData))
		return
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="localhost:8080"`)//localhost
	if strings.Count(r.URL.String()[len("/basic-auth/"):],"/") != 1 {
		http.Error(w,"Not Found",http.StatusNotFound)
		return
	}
	userAndpasswd := r.URL.String()[len("/basic-auth/"):]
	user := userAndpasswd[:strings.Index(userAndpasswd,"/")]
	pass := userAndpasswd[strings.Index(userAndpasswd,"/")+1:]
	if !check(user,pass,r){
		http.Error(w,"Unauthorised Attempt",http.StatusUnauthorized)
		return
	}

	w.Write(makeJSONresponse(jsonData))
	log.Println("User logged in:",user)
	
}
//hiddenBasicAuthHandler handles a GET request and sends a response in JSON format.
//It recieves password and username from client and checks that it is valid or not.
func hiddenBasicAuthHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="localhost:8080"`)	
	if len(r.URL.String())==12 || strings.Count(r.URL.String()[len("/basic-auth/"):],"/") != 1 {
		http.Error(w,"Not Found",http.StatusNotFound)
		return
	}
	userAndpasswd := r.URL.String()[len("/basic-auth/"):]
	user := userAndpasswd[:strings.Index(userAndpasswd,"/")]
	pass := userAndpasswd[strings.Index(userAndpasswd,"/")+1:]
	if !check(user,pass,r){
		http.Error(w,"Unauthorised Attempt",http.StatusUnauthorized)
		return
	} 
	http.Error(w,"Not Found",http.StatusNotFound)		
}

//streamHandler handles a GET request and sends a response in JSON format that contains url,args,headers,IP of the coming request.
//It sends response n times.
func streamHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	var n int
	nparam, err := strconv.ParseInt(r.URL.String()[len("/stream/"):],10,64)
	switch{
	case err != nil:
		n = 20
	case int(nparam)>100:
		n = 100
	case int(nparam)<=100:
		n = int(nparam)
	}
	jsonData := jsonMap{}
	jsonData = getAllJSONdata(r,"url","args","headers","origin")
	for i:=0;i<n;i++{
		jsonResp,_:= json.Marshal(jsonData)
		w.Write(jsonResp)
		w.Write([]byte("\n"))
	}
}

//delayHandler handles a GET request and sends a response in JSON format that contains args,data,files,form,headers,IP,url of the coming request.
//It sends response with a delayed time according to given n.
func delayHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	var n int
	nparam, err := strconv.ParseInt(r.URL.String()[len("/delay/"):],10,64)
	switch{
	case err != nil:
		n = 3
	case int(nparam)>10:
		n = 10
	case int(nparam)<=10:
		n = int(nparam)
	}
	time.Sleep(time.Second * time.Duration(n))
	jsonData := jsonMap{}
	jsonData = getAllJSONdata(r,"args","data","files","form","headers","origin","url")
	w.Write(makeJSONresponse(jsonData))
}

//htmlHandler handles a GET request and sends a sample HTML template.
func htmlHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	templates.SampleTemplate.ExecuteTemplate(w,"sample",nil)
}

//robotsTextHandler handles a GET request and sends a message that contains some robots.txt rules.
func robotsTextHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	w.Write([]byte("User-agent: *\nDisallow: /deny"))

}

//denyHandler handles a GET request and sends a message which recites that denied by robots.txt rules.
func denyHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	w.Write([]byte(` 
	  .-''''''-.
        .' _      _ '.
       /   O      O   \\
      :                :
      |                |
      :       __       :
        '.         	 .'
          '-......-'
     YOU SHOULDN'T BE HERE`))
}

//imageHandler handles a GET request and redirects it to "https://httpbin.org/image". 
func imageHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	http.Redirect(w,r,"https://httpbin.org/image",http.StatusOK)
}

//pngHandler handles a GET request and redirects it to "https://httpbin.org/image/png". 
func pngHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	http.Redirect(w,r,"https://httpbin.org/image/png",http.StatusOK)
}

//jpegHandler handles a GET request and redirects it to "https://httpbin.org/image/jpeg". 
func jpegHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	http.Redirect(w,r,"https://httpbin.org/image/jpeg",http.StatusOK)
}

//webpHandler handles a GET request and redirects it to "https://httpbin.org/image/webp". 
func webpHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	http.Redirect(w,r,"https://httpbin.org/image/webp",http.StatusOK)
}

//svgHandler handles a GET request and redirects it to "https://httpbin.org/image/svg". 
func svgHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	http.Redirect(w,r,"https://httpbin.org/image/svg",http.StatusOK)
}

//formsHandler handles a GET request and a sends sample form template.
func formsHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	templates.FormsTemplate.ExecuteTemplate(w,"forms",nil)
}

//xmlHandler handles a GET request and sends sample XML template.
func xmlHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	w.Header().Set("Content-Type","application/xml")
	templates.XmlTemplate.ExecuteTemplate(w,"xml",nil)
}

//linkHandler handles a GET request and sends n numbers link.
func linkHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	var n int
	strArr := strings.Split(r.URL.String(),"/")
	nparam,err := strconv.ParseInt(strArr[2],10,64)
	switch{
	case err != nil:
		n = 10
	case int(nparam)>200:
		n = 200
	case int(nparam)<=200:
		n = int(nparam)
	}
	var html []string
	html = append(html,"<html><head><title>Links</title></head><body>")
	for i:=0;i<n;i++{
		html = append(html,fmt.Sprintf(` <a href=/links/%d/%d> %d </a> `,n,i,i))
	}
	html = append(html,"</body></html>")
	resp := strings.Join(html,"")
	w.Write([]byte(resp))
}

//cacheHandler handles a GET request and sends a response in JSON format that contains url,args,header,IP of the coming request.
//If "If-Modified-Since" or "If-None-Match" header is provided,it returns 304 status code.
func cacheHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	jsonData := getAllJSONdata(r,"url","args","headers","origin")
	//http date and uuid check
	if r.Header.Get("If-Modified-Since" ) == "" && r.Header.Get("If-None-Match") == ""{
		//uuid :=getAllJSONdata(r,"uuid")["uuid"]
		w.Header().Set("Last-Modified","")
		//w.Header().Set("ETag",strconv.Itoa([]byte(uuid.([]uint8))))
		w.Write(makeJSONresponse(jsonData))
	}else{
		w.WriteHeader(304)
	}
}

//cacheControlHandler handles a GET request and sends a response in JSON format that contains url,args,headers,IP of the coming request.
//It sets a Cache-Control header for n seconds.
func cacheControlHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	//url check
	strArr := strings.Split(r.URL.String(),"/")
	nparam,err := strconv.ParseInt(strArr[2],10,64)
	if err != nil {
		nparam = 60
	}
	jsonData := getAllJSONdata(r,"url","args","headers","origin")
	w.Header().Set("Cache-Control",fmt.Sprintf("public, max-age=%d",nparam))
	w.Write(makeJSONresponse(jsonData))
}

//bytesHandler handles a GET request and sends a response that contains bytes which are generated n times randomly.
func bytesHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Method Not Allowed",405)
		return	
	}
	var n int
	var byteArr []byte
	urlStrArr := strings.Split(r.URL.String(),"/")
	nparam, err := strconv.ParseInt(urlStrArr[2],10,64)
	switch{
	case err != nil:
		n = 1024
	case int(nparam)>100*1024:
		n = 100*1024
	case int(nparam)<=100*1024:
		n = int(nparam)
	}
	for i:=0;i<n;i++{
		byteArr = append(byteArr,byte(rand.Int()))
	}
	w.Header().Set("Content-Type","application/octet-stream")
	w.Write(byteArr)
}