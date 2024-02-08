# Bell

A microservice for SMS and MMS messaging over Gmail.

## Configuration

### Environment Setup

The following environment variables are configurable when launching Bell:

- **BELL_HOSTNAME=** hostname for bell (default "127.0.0.1")
- **BELL_PORT=** port for bell (default 8099)
- **BELL_EMAIL=** sender gmail address
- **BELL_EMAIL_PWD=** sender gmail password

## **Endpoints**

The following REST endpoints are available:

- `/sms/send` **[POST]** Send SMS to recipient

## **Setup**

### **Bare-Metal**

1. Install Golang(1.21) onto the machine.
2. Setup Firebase and Redis as specified [here](#configuration).
3. Build the application with `go build`.
4. Launch the application with the appropriate flags via the `bell` binary.

### **Docker**

1. Build the container `docker build -t bell:1.0 .` . Note, the [environment setup](#environment-setup)
options can be passed here as build args to make them built into
the image.
