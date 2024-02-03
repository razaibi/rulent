### Rulent

Rulent is a Golang based low latency events engine that uses a set of declarative rules and triggers actions (outcomes).

### Key Components

The main components of the event engine are as below.

#### The Web API

This is for ingesting events. This acts as the entry point for new events to be received and validation against rules.

#### Events

Events are specific, explicitly sent, pieces of information sent with the payload. This sets the foundation for the rules engine to process rules based on the event.

#### Rules

These are the rules that need to be applied on the JSON payload. This has support for dot notation. The rules can be combined using `rules-operator: "or"`. 

#### Outcomes

Once the payloads have been processed against rules, matched payloads can initiate actions. **Actions** are bundled together as outcomes.

Note: Actions are executed asynchronously.


#### Actions 

Actions that are part of outcomes can be extended by navigating to the logic folder and actions.go file. Currently, this has access to outcomes and the payloads that were received.


### Sample Clients.

Make a request to validate for an "or" based combination of rules.

```bash
curl --location --request POST 'http://localhost:8081/validate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "events" : ["purchase"],
    "person": {
        "name": "James",
        "lname": "Wright",
        "age": "39"
    }
}
'
```

Make a request to validate for an "and" based combination of rules.

```bash
curl --location --request POST 'http://localhost:8081/validate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "events" : ["click"],
    "person": {
        "name": "James",
        "age": "42"
    }
}
'
```

The sample **events.yaml** file is included in the project folder.

### Performance Tuning

Some obvious performance tuning approach would be the following.

- Adding a load balancer in front of the Web API.
- Run the service on infrastructure meant for higher in memory compute.
- Limit complexity of outcomes.

### Roadmap

- Add custom outcomes
    - REST API Call
- Add support for JSON based config.



