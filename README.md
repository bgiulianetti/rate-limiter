
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
- Rules and notifications are handled by two different services, and their persistences as well.
- Initial rules are obtained from a json file, and they are saved in the  rules memory repository and they are handled by the Rules Service.
- The Notifications can be stored in memory or in Redis. To configure this, an environment variable must be set in the [Makefile](https://github.com/bgiulianetti/rate-limiter/blob/main/makefile#L7)
- Rules can only be stored in memory, but the implementation can easily be adapted to be stored in Redis (or any other database)
- The API is prepared to handle multiple rules by notification type.

## Local Development Setup
- To run the API you will have to run this command first: ```make initialize```.<br />
This command will:
  - Install all of the dependencies
  - Create all of the mocks
  - Run all the tests
  - Run the API
- The API runs in the port ```5000``` by default, but you can change it [here](https://github.com/bgiulianetti/rate-limiter/blob/main/main.go#L12)
- The API by default uses in memory storage for Notifications, but you can change it to use Redis [here](https://github.com/bgiulianetti/rate-limiter/blob/main/makefile#L7)
- The Redis server is up and running, you can run your API locally, configure to use Redis and it will work properly.


