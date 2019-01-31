package main

// Main entry for program.
func main() {
	LogInf.Println("Starting web server...")
	// Starting the web service blocks the main thread until it
	// crashes or the prgram was signaled to stop from outside.
	// If starting the web server fails, it returns an error which
	// be output to stderr.
	LogFtl.Println("Failed starting web service: ", StartHTTPServer("8080"))
}
