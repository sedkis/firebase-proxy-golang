## Firebase Proxy Server

This is a proxy that will accept a user's email/password and forward it 
to the [Firebase Auth API](https://firebase.google.com/docs/reference/rest/auth#section-sign-in-email-password).

## Running
Run from source, or build the Docker image and add your API key as an env variable.

### From docker
1. First, we build the image

```
$ docker build -t firebaseproxyserver .
``` 

2. Run the image
```
$ docker run -p 9001:9001 -e GOOGLE_API_KEY=123 firebaseproxyserver
```

3. Make an API request
```
$ curl 'http://localhost:9001?email=hello@world.com&password=123'

{
  "email": "hello@world.com",
  "idToken": "eyJhbGc...",
  ...
}
```