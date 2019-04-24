package wmata

import "os"

// wmata constants
const baseURL = "https://api.wmata.com"

var apiKey = os.Getenv("WMATA_API_KEY")
