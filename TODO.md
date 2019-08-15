# app engine API

WIP

# access to GKE

DONE

# dialogflow

https://godoc.org/google.golang.org/api/dialogflow/v2#GoogleCloudDialogflowV2WebhookRequest

commands logic. store context

GetSession(sessionID) // create existing one or get by current id
session:
    id
    context: welcome|

action: scale
resource: deployment
name: redis

WIP

1. basic intents
2. dialogflow api
3. full conversation model
4. export as zip
5. error handling

Intents:
    welcome
    scale_req
    do_scale

Conversation:

bot: Hi
1. waiting for user command
user: no/quit - end conversation
user: command
    if error: say error and go to 1
    if not recognized: say error and go to 1
    additional command question
        2. waiting for user reply
        if error: say error and go to 1
        if not recognized: say error and go to 2
        success: say result and go to 1
        if quit: "Can I help you with something else"? go to 1


Global commands:
- quit/exit/stop - ends conversation