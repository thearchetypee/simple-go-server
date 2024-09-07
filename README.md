# Simple Go Server

Simple Go Server is a RESTful API built with Go, designed to fetch and filter data from MongoDB and provide in-memory data storage functionality.

## Features

- Fetch and filter data from MongoDB based on date range and count
- In-memory key-value storage with RESTful endpoints
- High performance and scalability
- Comprehensive error handling and logging

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16+
- MongoDB

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/thearchetypee/simple-go-server.git
   ```

2. Navigate to the project directory:
   ```
   cd simple-go-server
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

4. Set up your environment variables (see "Setting up .env file" section below)

5. Run the application:
   ```
   make
   ```

The server should now be running on `http://localhost:3000` (or the port specified in your .env file).

### Setting up .env file

This project contains a `.env_sample` file with two variables: `SERVER_PORT` and `MONGO_URI`. Follow these steps:

1. Rename `.env_sample` to `.env`
2. Add your MongoDB URI. If you've set up MongoDB locally, it will look like this: `mongodb://127.0.0.1:27017`
3. Adjust the `SERVER_PORT` if needed (default is 3000)

### Setting Up MongoDB

1. Follow this guide to set up a local MongoDB instance: [Setting up MongoDB on macOS](https://www.prisma.io/dataguide/mongodb/setting-up-a-local-mongodb-database#setting-up-mongodb-on-macos)

2. This project includes a `records.json` file containing sample data. To populate the database:
   - Ensure your MongoDB URI is correctly set in the .env file
   - Locate the `insert()` function in the project (you may need to add a call to this function in your main.go or a setup script)
   - Run the function to insert the sample data into your MongoDB instance

## Usage

### Fetching Data from MongoDB

Send a POST request to `/mongo` with the following JSON payload:

```json
{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "minCount": 2700,
    "maxCount": 3000
}
```

### In-Memory Data Storage

#### Storing Data

Send a POST request to `/in-memory` with the following JSON payload:

```json
{
    "key": "2",
    "value": "here we are 3"
}
```

#### Retrieving Data

Send a GET request to `/in-memory?key=2`

## API Reference

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/mongo` | POST | Fetch and filter data from MongoDB |
| `/in-memory` | POST | Store data in-memory |
| `/in-memory` | GET | Retrieve data from in-memory storage |
