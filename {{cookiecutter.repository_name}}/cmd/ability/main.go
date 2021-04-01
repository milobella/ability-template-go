package main

import (
	"github.com/milobella/ability-sdk-go/pkg/ability"
)

// fun main()
func main() {
	// Read configuration
	conf := ability.ReadConfiguration()

	// Initialize server
	server := ability.NewServer("{{cookiecutter.project_title}}", conf.Server.Port)
	server.RegisterIntentRule("DEFAULT_INTENT", DefaultIntentHandler)
	server.Serve()
}

func DefaultIntentHandler(_ *ability.Request, resp *ability.Response) {
	// Build the NLG answer
	resp.Nlg.Sentence = "I don't know how to do this for the moment."
}
