package main
/*
*Author William Kwesidev
*Email willzako@aol.com
*/



import("net/http"
       "io"
       "fmt"
       "time"
       "flag"
       "os"
       "io/ioutil"
       "log"
)


func welcome(res http.ResponseWriter,req *http.Request){

	res.Header().Set(
		"Conten-Type",
		"text/html",
		)

	io.WriteString(res,
		          `<!doctype html>
	<html>
	   <head>
		    <meta charset=utf-8>
		    <title>Go lang WebApp</title>
		 </head>
		 <body bgColor=yellow>
		 	<h1>GOLANG IS WONDRFUL</h1>
		 	<h3> Powered By </h3>
			<!--  Base64 encode image -->
			<img src="data:image/png;base64,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
QmCC
" />
		 </body>
	</html>`,
		)

   fmt.Fprint(os.Stdout,time.Now()," /welcome")
}
func main(){

	port := flag.String("port","8080","Port to listen to")
	//configure http server
     server := &http.Server{
    		Addr:fmt.Sprintf(":%s",*port),
    		Handler: nil,
    		ReadTimeout: 10 * time.Second,
    		WriteTimeout: 10 * time.Second,
    		MaxHeaderBytes: 1 << 9,
    } 


	http.HandleFunc("/",welcome)
	//gets latest news from cnn
	http.HandleFunc("/news",func (res http.ResponseWriter, req *http.Request){

			res.Header().Set("Conten-Type","text/html")

			web,err := http.Get("http://rss.cnn.com/rss/edition_world.rss")
			if err != nil {

				return
			}

			defer web.Body.Close()
			data ,err := ioutil.ReadAll(web.Body)

			if err != nil{
				log.Fatal("Failed to download data")
				return 
			}

			io.WriteString(res,string(data))
			fmt.Fprint(os.Stdout,time.Now(), " /news")


		})
	fmt.Println(time.Now())
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
