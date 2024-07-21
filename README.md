# ShortNGo

ShortNGo is a URL shortening service built with Go and powered by HTMX. It uses [BadgerDB](https://github.com/dgraph-io/badger) as the underlying storage mechanism to store and retrieve shortened URLs.

## Features

- Shorten long URLs.
- Redirect to the original URL using the shortened URL.
- Uses a hash function to generate unique shortened URLs.
- In-memory database using BadgerDB for fast performance.

## Prerequisites

- Go 1.16 or higher
- Internet connection to fetch dependencies

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/shortngo.git
cd shortngo
```

2. Install dependencies:

```bash
go mod tidy
```

## Usage

1. Run the application:

```bash
go run .
```

2. Open your web browser and go to `http://localhost:8080`.

### Endpoints

- **POST /process**

  This endpoint accepts a URL via a form and generates a shortened version.

- **GET /go/{hash}**

  This endpoint redirects to the original URL associated with the given hash.

  Example:

  ```bash
  curl http://localhost:8080/go/c19fcdc5
  ```

## Project Structure

- `main.go` - Entry point of the application, handles HTTP routing and request processing.
- `database/*` - Contains the database implementations and factory.

## Testing

To run the tests, use the following command:

```bash
go test ./...
```

This will execute the tests defined in `database/badgerim_test.go` to ensure the database implementation is working as expected.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

## Acknowledgements

- [BadgerDB](https://github.com/dgraph-io/badger) for the in-memory database.
- [HTMX](https://htmx.org/) for handling AJAX requests in a simpler way.
- [Murmur3](https://github.com/spaolacci/murmur3) for the hashing function.

---

Feel free to customize this README to better fit your project and personal style.
