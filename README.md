# Go REST API Template

A streamlined Go-based REST API template integrating Echo, SQLC, and clean architecture principles. This template serves as a robust starting point for building high-performance backend services, emphasizing clean code structure, ease of configuration, security best practices, and efficient database interactions.

## Quick Start

**Clone and Install:**

```bash
git clone github.com/HoneySinghDev/go-templ-htmx-template.git
cd go-templ-htmx-template
task install
```

**Launch Development Server:**

```bash
task dev
```

## Configuration

**Environment Variables:**

Create a `.env` file based on `.env.example`. Configure database and other service settings to match your development or production environments.

**PKL Configuration:**

Settings are managed through `pkl/app.config`, with environment-specific settings located in `pkl/local` and `pkl/prod`.

## Building for Production

```bash
task build
```

Compiles the Go binary. Deploy `bin/app` directly to your production server.

## Project Structure

- `cmd`: Entry points for application execution.
- `internal`: Application-specific logic (router, handlers, services, repositories).
- `pkg`: Shared packages and utilities.
- `pkl`: Environment-specific configuration.

## Logging & Monitoring

The project uses zerolog for structured and performant logging. Adjust the logging level and formatting via environment variables.

## Database Integration

Utilizes SQLC for type-safe, efficient database interactions. Ensure your database schema and queries are updated in the `sqlc` configuration files.

## Graceful Shutdown

The server includes built-in graceful shutdown handling, ensuring reliability and data integrity during server restarts and deployments.

## Contributing

Contributions are welcome! Fork this repository, implement your changes, and submit a pull request.

## License

MIT License. See `LICENSE` for full details.

---

Simplify your Go backend development with this clean and efficient REST API template, crafted for productivity and scalability.
