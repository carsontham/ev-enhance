## EV Enhance 

This project is built during a hackathon - Hackomania 2024.

EV Enhance aims to create an API-centric app that allows for a one-stop platform for all the different major EV companies. The MVP aims to onboard all EV Partners and interface them to one platform.

The backend service calls a Mock Postman Server and is meant to be paired together with [`ev-enhance-ios`](https://github.com/carsontham/ev-enhance-ios), an iOS application developed in Swift.

## Steps to run backend service

1. Ensure Golang is installed.

https://go.dev

2.  Ensure golang-migrate is installed.

```
brew install golang-migrate
```

3. Run the following cmd for setup

```
make start-and-migrate
```

4. To run the backend service at port :3000

```
make server
```

## Steps to run iOS application
1. Ensure Xcode is installed from App Store
2. Run the application on an iPhone 12 Emulator
3. Ensure backend service is up to test the full MVP.
