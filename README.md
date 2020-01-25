# :calendar: :white_check_mark: Daily Tasks App

This is a daily-task list application.

**Server: Golang  
Client: React, semantic-ui-react  
Database: Local MongoDB**

---

# :pen: Application Requirement

### golang server requirement

1. golang https://golang.org/dl/
2. gorilla/mux library for router `go get -u github.com/gorilla/mux`
3. mongo-driver library to connect with mongoDB `go get go.mongodb.org/mongo-driver`

### react client

From the Application directory

`create-react-app client`

# :vertical_traffic_light: Start the application

1. Make sure your mongoDB is started
2. From server directory, open a terminal and run
   `go run main.go`
3. From client directory,  
   a. install all the dependencies using `npm install`  
   b. start client `npm start`

# :walking: Walk through the application

Open application at http://localhost:3000

#### :books: For begginers getting started with GoLang via examples: go-tutorial directory

# References

https://godoc.org/go.mongodb.org/mongo-driver/mongo  
https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial  
https://vkt.sh/go-mongodb-driver-cookbook/
