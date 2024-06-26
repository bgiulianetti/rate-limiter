
# [Backend] Rate-Limited Notification Service
## The task
We have a Notification system that sends out email notifications of various types (status update, daily news, project invitations, etc). We need to protect recipients from getting too many emails, either due to system errors or due to abuse, so let’s limit the number of emails sent to them by implementing a rate-limited version of NotificationService.

The system must reject requests that are over the limit.

Some sample notification types and rate limit rules, e.g.:
- Status: not more than 2 per minute for each recipient
- News: not more than 1 per day for each recipient
- Marketing: not more than 3 per hour for each recipient

Etc. these are just samples, the system might have several rate limit rules!

## Decisions made
### Development
- The project consists of a REST API developed in Golang with [Gin](https://github.com/gin-gonic/gin).
- The Storage is handled in memory and with [Redis](https://redis.io/)
- Interface mocks are handled with [Moq](https://github.com/matryer/moq)
### Business Logic
- Rules and notifications are handled by two different services, and their persistence as well.
- Initial rules are obtained from a json file, and they are saved in the  rules memory repository and handled by the Rules Service.
- The Notifications can be stored in memory or in Redis. To configure this, an environment variable must be set in the [Makefile](https://github.com/bgiulianetti/rate-limiter/blob/main/makefile#L7)
- Rules can only be stored in memory, but the implementation can easily be adapted to be stored in Redis (or any other database)
- The API is prepared to handle multiple rules by notification type.
- If a notification type has no rule, it is possible to send as many notifications as desired.

## Local Development Setup
- To run the API for the first time, it is mandatory to run this command first:
  ```
  make initialize
  ```
This command will:
  - Install all of the dependencies
  - Create all of the mocks
  - Run all the tests
  - Run the API
- if it is not the first time, it is possible to run the API with any of these commands:
  - ```make all``` This will run all of the tests and run the API.
  - ```make run``` This will run the API.
- The API runs on the port ```5000``` by default, but it can be changed [here](https://github.com/bgiulianetti/rate-limiter/blob/main/main.go#L12)
- The API by default uses in memory storage for Notifications, but it can be changed to use Redis [here](https://github.com/bgiulianetti/rate-limiter/blob/main/makefile#L7)
- The Redis server is up and running, the API can run locally, it can be configured to use Redis and it will work properly.

## Endpoint
### Request
```
POST /notifications/:type/users/:user
```

### Responses

OK - HTTP Status code: 200
```
{
    "message": "notification sent",
    "status": "success"
}
```

Too many requests - HTTP status code: 429
```
{
    "message": "message limit exceeded",
    "error": "rate limit exceeded",
    "status": 429
}
```


