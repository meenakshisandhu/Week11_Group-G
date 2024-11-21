# **Go Web Server for Static File Serving**
 
This project implements a Go-based web server capable of serving static files from a specified directory. The server is designed to handle various HTTP requests and serve files appropriately. It includes comprehensive unit tests to ensure the server functions as expected. 
 
---
 
## **Features**
- Serve static files (HTML, CSS, JavaScript) from a specified directory (./Static).
- Handle requests for non-existent files and return appropriate HTTP status codes.
- Unit tests to ensure the server's correctness in handling file requests, non-existent files, and content types.
 
---
## **Project Structure **
``` bash
/project
│
├── /Static              # Directory containing static files (HTML, CSS, JS, TEXT)
│   ├── index.html       # Example HTML file
│   ├── style.css        # Example CSS file
│   └── script.js        # Example JS file
│   └── testfile.txt     # Example Text file
|
├── server_test.go       # Test file for the Go web server
│
├── main.go              # Go application for serving static files
├── README.md            # Project documentation
└── go.mod               # Go module file
```

## **Commands**
Start the server:
``` bash
go run main.go
```
Access the application in your browser at:
``` bash
http://localhost:8080
```

![Web Server](https://github.com/meenakshisandhu/Week11_Group-G/blob/master/images/Screenshot%202024-11-20%20212021.png)


Run the test suite:
``` bash
go test -v
```

![Testing](https://github.com/meenakshisandhu/Week11_Group-G/blob/master/images/Screenshot%202024-11-20%20212132.png)


Testing includes:
- Validating HTTP status codes for existing and non-existent files.
- Verifying MIME types for HTML, CSS, and JavaScript files.
- Ensuring the default file (index.html) serves correctly.
 
# **Test Scenarios**
#### Serving Existing Files:
- Verifies that files present in the server's directories are served correctly.
- Ensures the server responds with:
- HTTP Status Code: 200 OK.
Handling Non-Existent Files:
- Tests requests for files that do not exist on the server.
#### Confirms the server returns:
- HTTP Status Code: 404 Not Found.
#### MIME Type Validation:
- Ensures that files are served with the correct MIME type based on their extensions.
#### MIME types validated:
- HTML → Static/index.html
- CSS → Static/style.css
- JavaScript → Static/script.js
#### Default File Handling:
- Validates that the server serves the default index.html file when the root path (/) is accessed.
#### Confirms the server responds with:
- HTTP Status Code: 200 OK.
