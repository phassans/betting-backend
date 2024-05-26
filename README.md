### Betting Backend

---

### Introduction

This project involves running server APIs, setting up a PostgreSQL database, processing smart contract events, saving these events into a database, and providing APIs about events and bets.

---

### 1. Running Server APIs

#### Prerequisites

- Golang (version 1.18 or above)
- PostgreSQL

#### Setup

1. **Clone the repository**

   ```sh
   git clone https://github.com/your-repo/project-name.git
   cd project-name
   ```

2. **Install dependencies**

   ```sh
   go mod tidy
   ```

3. **Run the server**

   ```sh
   go run main.go
   ```

4. **Environment Variables**

   Ensure the following environment variables are set in your system or in a `.env` file:

   ```sh
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=youruser
   DB_PASSWORD=yourpassword
   DB_NAME=yourdbname
   SERVER_PORT=8080
   ```

---

### 2. Setting Up PostgreSQL

#### Prerequisites

- PostgreSQL

#### Setup

1. **Install PostgreSQL**

   Follow the installation instructions for your operating system from the [official PostgreSQL documentation](https://www.postgresql.org/docs/).

2. **Start PostgreSQL**

   Ensure PostgreSQL is running.

3. **Create a Database**

   ```sh
   psql -U postgres
   CREATE DATABASE yourdbname;
   CREATE USER youruser WITH ENCRYPTED PASSWORD 'yourpassword';
   GRANT ALL PRIVILEGES ON DATABASE yourdbname TO youruser;
   ```

4. **Connect to PostgreSQL**

   ```sh
   psql -h localhost -U youruser -d yourdbname
   ```

---

### 3. Processing Smart Contract Events

The file `read-events.go` handles the processing of smart contract events. Key steps include:

1. **Connecting to the Blockchain Network**

   Ensure you have the correct RPC endpoint and contract address.

   ```go
   client, err := ethclient.Dial("https://your-ethereum-node")
   if err != nil {
       log.Fatal(err)
   }
   ```

2. **Listening to Events**

   Utilize the `go-ethereum` package to set up event subscriptions.

   ```go
   logs := make(chan types.Log)
   sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
   if err != nil {
       log.Fatal(err)
   }

   for {
       select {
       case err := <-sub.Err():
           log.Fatal(err)
       case vLog := <-logs:
           // Process event
       }
   }
   ```

3. **Processing Events**

   Implement logic to process and handle the events.

   ```go
   for {
       select {
       case err := <-sub.Err():
           log.Fatal(err)
       case vLog := <-logs:
           eventName := vLog.Topics[0].Hex()
           eventData := vLog.Data
           // Save event to DB
       }
   }
   ```

---

### 4. Saving Smart Contract Events into DB

The processed events are saved into a PostgreSQL database. The schema for the tables is defined below:

#### Database Schema

```sql
-- Table: events
CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    tx_hash VARCHAR(255),
    block_number BIGINT,
    block_hash VARCHAR(255),
    input_data BYTEA,
    bet_id BIGINT,
    event_type VARCHAR(255),
    contract_address VARCHAR(255),
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: bets
CREATE TABLE bets (
    id SERIAL PRIMARY KEY,
    bet_id BIGINT UNIQUE NOT NULL,
    user_a VARCHAR(255),
    user_b VARCHAR(255),
    amount DOUBLE PRECISION,
    winner VARCHAR(255),
    reward DOUBLE PRECISION,
    is_long BOOLEAN,
    create_time TIMESTAMP,
    expire_time TIMESTAMP,
    closing_time TIMESTAMP,
    opening_price DOUBLE PRECISION,
    closing_price DOUBLE PRECISION,
    bet_status INT
);
```

#### Inserting Events

Use the `GORM` package to interact with the database.

```go
dsn := "host=localhost user=youruser password=yourpassword dbname=yourdbname port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
}
```

---

### 5. APIs About Events and Bets

The API endpoints are defined in the `api.go` file. Key endpoints include:

## Base URL

The base URL for all endpoints is: `http://<your-domain>/`

## Endpoints

### Bets

#### Get All Bets

- **URL**: `/bets`
- **Method**: `GET`
- **Description**: Retrieves a list of all bets.
- **Response**:
    - `200 OK`: Returns an array of bet objects.

#### Get Bet by ID

- **URL**: `/bets/:id`
- **Method**: `GET`
- **Description**: Retrieves a bet by its ID.
- **URL Parameters**:
    - `id` (uint): The ID of the bet.
- **Response**:
    - `200 OK`: Returns the bet object.
    - `404 Not Found`: Bet not found.

#### Create a New Bet

- **URL**: `/bets`
- **Method**: `POST`
- **Description**: Creates a new bet.
- **Request Body**: JSON object representing the bet.
    - Example:
      ```json
      {
        "BetID": 123456,
        "UserA": "Alice",
        "UserB": "Bob",
        "Amount": 50.0,
        "Winner": "",
        "Reward": 100.0,
        "IsLong": true,
        "CreateTime": "2024-05-26T12:00:00Z",
        "ExpireTime": "2024-06-26T12:00:00Z",
        "ClosingTime": "2024-07-26T12:00:00Z",
        "OpeningPrice": 200.0,
        "ClosingPrice": 250.0,
        "BetStatus": 0
      }
      ```
- **Response**:
    - `200 OK`: Returns the created bet object.
    - `400 Bad Request`: Invalid request body.

#### Update a Bet by ID

- **URL**: `/bets/:id`
- **Method**: `PUT`
- **Description**: Updates a bet by its ID.
- **URL Parameters**:
    - `id` (uint): The ID of the bet.
- **Request Body**: JSON object representing the updated bet.
    - Example:
      ```json
      {
        "UserA": "Alice",
        "UserB": "Bob",
        "Amount": 60.0,
        "Winner": "Alice",
        "Reward": 120.0,
        "IsLong": false,
        "CreateTime": "2024-05-26T12:00:00Z",
        "ExpireTime": "2024-06-26T12:00:00Z",
        "ClosingTime": "2024-07-26T12:00:00Z",
        "OpeningPrice": 200.0,
        "ClosingPrice": 250.0,
        "BetStatus": 1
      }
      ```
- **Response**:
    - `200 OK`: Returns the updated bet object.
    - `400 Bad Request`: Invalid request body or bet ID.
    - `404 Not Found`: Bet not found.
    - `500 Internal Server Error`: Server error.

### Events

#### Get All Events

- **URL**: `/events`
- **Method**: `GET`
- **Description**: Retrieves a list of all events.
- **Response**:
    - `200 OK`: Returns an array of event objects.

#### Get Event by ID

- **URL**: `/events/:id`
- **Method**: `GET`
- **Description**: Retrieves an event by its ID.
- **URL Parameters**:
    - `id` (uint): The ID of the event.
- **Response**:
    - `200 OK`: Returns the event object.
    - `404 Not Found`: Event not found.

---

Ensure to replace `<your-domain>` with the actual domain where the APIs are hosted.
