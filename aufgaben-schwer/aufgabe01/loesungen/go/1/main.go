package main

func main() {
	LogInf.Println("Starting web server...")
	LogFtl.Println("Failed starting web service: ", StartHTTPServer("8080"))
}
