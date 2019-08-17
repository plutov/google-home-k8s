package dialogflow

// Dialogflow V2 webhook request and response types

// Request struct
type Request struct {
	QueryResult struct {
		Parameters map[string]interface{} `json:"parameters"`
		Action     string                 `json:"action"`
	} `json:"queryResult"`
	OriginalRequest OriginalRequest `json:"originalDetectIntentRequest"`
	Session         string          `json:"session"`
}

// OriginalRequest struct
type OriginalRequest struct {
	Payload OriginalRequestPayload `json:"payload"`
}

// OriginalRequestPayload struct
type OriginalRequestPayload struct {
	Device OriginalRequestDevice `json:"device"`
	User   OriginalRequestUser   `json:"user"`
}

// OriginalRequestDevice struct
type OriginalRequestDevice struct {
	Location OriginalRequestLocation `json:"location"`
}

// OriginalRequestUser struct
type OriginalRequestUser struct {
	Permissions []string `json:"permissions"`
}

// OriginalRequestLocation struct
type OriginalRequestLocation struct {
	Coordinates OriginalRequestCoordinates `json:"coordinates"`
}

// OriginalRequestCoordinates struct
type OriginalRequestCoordinates struct {
	Lat  float32 `json:"latitude"`
	Long float32 `json:"longitude"`
}

// Response struct
type Response struct {
	Payload ResponsePayload `json:"payload"`
}

// ResponsePayload struct
type ResponsePayload struct {
	Google ResponseGoogle `json:"google"`
}

// ResponseGoogle struct
type ResponseGoogle struct {
	ExpectUserResponse bool                  `json:"expectUserResponse"`
	RichResponse       RichResponse          `json:"richResponse"`
	SystemIntent       *ResponseSystemIntent `json:"systemIntent,omitempty"`
}

// RichResponse struct
type RichResponse struct {
	Items []Item `json:"items"`
}

// Item struct
type Item struct {
	SimpleResponse SimpleResponse `json:"simpleResponse"`
}

// SimpleResponse struct
type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
}

// ResponseSystemIntent struct
type ResponseSystemIntent struct {
	Intent string                   `json:"intent"`
	Data   ResponseSystemIntentData `json:"data"`
}

// ResponseSystemIntentData struct
type ResponseSystemIntentData struct {
	Type        string   `json:"@type"`
	OptContext  string   `json:"optContext"`
	Permissions []string `json:"permissions"`
}
