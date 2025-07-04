# 🔗 Shorten URL with Fiber & Redis

![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)

A simple, fast, and efficient URL shortener API built with the powerful combination of [Go](https://golang.org/), [Fiber](https://gofiber.io/), and [Redis](https://redis.io/).

This project provides a minimalistic API to shorten long URLs and redirect users to their original destinations using the generated short codes. It's designed for easy setup and deployment, thanks to Docker Compose support.

---

## ✨ Features

-   **Simple API:** Shorten any valid URL with a single, straightforward API call.
-   **Blazing Fast Redirection:** Leverages Redis as a high-speed backend for quick lookups and redirects.
-   **Dockerized:** Includes Docker Compose support for a hassle-free, one-command setup.
-   **API Versioning:** Basic versioning in the API path (`/api/v1`) for future-proofing.

---

## 🚀 Getting Started

You can get the project up and running in minutes using either Docker or by running it locally.

### Prerequisites

-   [Go](https://golang.org/) (v1.18 or higher)
-   [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/) (Recommended for easy setup)
-   [Redis](https://redis.io/) (if running locally without Docker)

---

### 🐳 Running with Docker Compose (Recommended)

This is the easiest way to get started.

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/shorten-url-fiber.git
    cd shorten-url-fiber
    ```

2.  **Build and start the containers:**
    ```sh
    docker-compose up --build
    ```

3.  The API will be running and accessible at `http://localhost:3000`.

### 💻 Running Locally

If you prefer to run the application on your local machine without Docker containers:

1.  **Start a Redis instance:**
    You can use Docker to easily start a Redis server.
    ```sh
    docker run --name my-redis -p 6379:6379 -d redis
    ```

2.  **Run the Go application:**
    From the project's root directory, run:
    ```sh
    go run main.go
    ```

3.  The API will be running and accessible at `http://localhost:3000`.

---

## 🛠️ API Endpoints

### Shorten a URL

Creates a new short URL for a given long URL.

-   **Endpoint:** `POST /api/v1`
-   **Headers:** `Content-Type: application/json`
-   **Request Body:**
    ```json
    {
      "url": "https://www.very-long-and-complex-url-to-shorten.com/with/some/path"
    }
    ```
-   **Success Response (200 OK):**
    ```json
    {
      "short_url": "http://localhost:3000/jA8sBf"
    }
    ```

### Redirect to Original URL

Redirects the user to the original long URL.

-   **Endpoint:** `GET /:shortCode`
-   **Example:** Navigating to `http://localhost:3000/jA8sBf` in your browser will redirect you to the original long URL.

---

## ⚙️ Configuration

The application can be configured using the following environment variables:

| Variable      | Description                  | Default          |
| :------------ | :--------------------------- | :--------------- |
| `REDIS_ADDR`  | The address of the Redis server. | `localhost:6379` |
| `SERVER_PORT` | The port for the API server. | `3000`           |

---

## 📄 License

This project is distributed under the **MIT License**. See the `LICENSE` file for more information.
