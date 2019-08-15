package dialogflow

// GenerateWebhookResponse .
func GenerateWebhookResponse(expectUserResponse bool, msg string) Response {
	return Response{
		Payload: ResponsePayload{
			Google: ResponseGoogle{
				ExpectUserResponse: expectUserResponse,
				RichResponse: RichResponse{
					Items: []Item{
						Item{
							SimpleResponse: SimpleResponse{
								TextToSpeech: msg,
							},
						},
					},
				},
			},
		},
	}
}
