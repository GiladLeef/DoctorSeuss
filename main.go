package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/GiladLeef/DoctorSeuss/locales"
	"github.com/GiladLeef/DoctorSeuss/training"

	"github.com/GiladLeef/DoctorSeuss/dashboard"

	"github.com/GiladLeef/DoctorSeuss/util"

	"github.com/gookit/color"

	"github.com/GiladLeef/DoctorSeuss/network"

	"github.com/GiladLeef/DoctorSeuss/server"
)

var neuralNetworks = map[string]network.Network{}

func main() {
	port := flag.String("port", "8080", "The port for the API and WebSocket.")
	localesFlag := flag.String("re-train", "", "The locale(s) to re-train.")
	flag.Parse()

	// If the locales flag isn't empty then retrain the given models
	if *localesFlag != "" {
		reTrainModels(*localesFlag)
	}

	// Print the Doctor Seuss ascii text
	DoctorSeussASCII := string(util.ReadFile("res/DoctorSeuss-ascii.txt"))
	fmt.Println(color.FgLightGreen.Render(DoctorSeussASCII))

	// Create the authentication token
	dashboard.Authenticate()

	for _, locale := range locales.Locales {
		util.SerializeMessages(locale.Tag)

		neuralNetworks[locale.Tag] = training.CreateNeuralNetwork(
			locale.Tag,
			false,
		)
	}

	// Get port from environment variables if there is
	if os.Getenv("PORT") != "" {
		*port = os.Getenv("PORT")
	}

	// Serves the server
	server.Serve(neuralNetworks, *port)
}

// reTrainModels retrain the given locales
func reTrainModels(localesFlag string) {
	// Iterate locales by separating them by comma
	for _, localeFlag := range strings.Split(localesFlag, ",") {
		path := fmt.Sprintf("english/training.json", localeFlag)
		err := os.Remove(path)

		if err != nil {
			fmt.Printf("Cannot re-train %s model.", localeFlag)
			return
		}
	}
}
