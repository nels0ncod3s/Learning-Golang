# 🚀 30-Day Go (Golang) Roadmap: Zero to Job-Ready

**Commitment:** 4 hours/day | **Total:** 120 hours | **Goal:** Build production-grade Go backends, pass junior/mid-level interviews

**Daily Rhythm:**
- **90 min** — Study & read documentation
- **120 min** — Hands-on coding (the only way to learn Go)
- **30 min** — Review, refactor yesterday's code, write notes

> **Core Philosophy:** *Go is not learned by watching — it is learned by building.* Every day ends with working code committed to GitHub.

---

## 📅 WEEK 1: Language Foundations (Days 1–7)
**Theme:** *"Master the grammar before writing poetry"*

### Day 1 — Setup, Tooling & First Programs
**Study (90 min):**
- Install Go 1.24+ (use `go.dev/dl`)
- Configure VS Code / GoLand with `gopls`, `gofumpt`, `golangci-lint`
- Understand `GOPATH` vs Go Modules (`go mod init`, `go mod tidy`)
- `go run`, `go build`, `go test`, `go fmt`, `go vet`
- Basic syntax: variables (`var`, `:=`), constants, `iota`, basic types, zero values

**Build (120 min):**
- Temperature converter (Celsius ↔ Fahrenheit ↔ Kelvin)
- Simple CLI calculator with `flag` package for arguments
- **Commit:** to GitHub

**Review (30 min):**
- Read [Effective Go](https://go.dev/doc/effective_go) — first 3 sections
- Set up daily Git commit habit

**🎯 Outcome:** Working Go environment, first committed project, comfort with basic syntax.

---

### Day 2 — Control Flow & Defer
**Study (90 min):**
- `if/else` (note: no parentheses around conditions)
- `switch` (expression switches, type switches — preview only)
- `for` (Go's only loop): C-style, condition-only, `range`
- `break`, `continue`, labeled breaks
- `defer`: LIFO execution, common patterns (resource cleanup)

**Build (120 min):**
- FizzBuzz with multiple implementations
- Number guessing game with difficulty levels
- File reader that uses `defer` to close files

**Review (30 min):**
- Experiment: What happens with multiple `defer` statements? (LIFO proof)

**🎯 Outcome:** Comfortable with all control structures; understand `defer` as Go's RAII.

---

### Day 3 — Functions: Go's Workhorse
**Study (90 min):**
- Function declarations, multiple return values
- Named return values (when to use, when to avoid)
- Variadic functions (`...args`)
- Closures (function values capturing outer scope)
- `panic` / `recover` (know they exist, avoid using them)

**Build (120 min):**
- String utility library: `Reverse`, `IsPalindrome`, `WordCount`, `Truncate`
- Closure-based counter factory
- Math operations with variadic params

**Review (30 min):**
- Compare named vs unnamed returns for readability

**🎯 Outcome:** Can write idiomatic Go functions; understand when to use closures.

---

### Day 4 — Arrays, Slices & Maps
**Study (90 min):**
- Arrays (fixed size, value type) vs Slices (dynamic, reference type)
- Slice internals: backing array, `len`, `cap`, `append`, `copy`
- Slice expressions, slicing beyond capacity
- Maps: creation, `delete`, zero-value check (`value, ok`)
- Iteration with `range` (index/value, key/value)

**Build (120 min):**
- Word frequency counter from text file
- Slice manipulation drills: reverse, filter, deduplicate
- Contact book CRUD using maps

**Review (30 min):**
- Visualize: Draw backing arrays for slice operations
- Benchmark: `append` vs pre-allocated slices

**🎯 Outcome:** Deep understanding of slice mechanics; can build data-processing tools.

---

### Day 5 — Structs, Methods & Pointers
**Study (90 min):**
- Structs, embedded structs (composition over inheritance)
- Methods: value receivers vs pointer receivers
- When to use pointers (large structs, mutation needed)
- Zero values and initialization
- `new()` vs composite literals

**Build (120 min):**
- `BankAccount` struct with `Deposit`, `Withdraw`, `Balance` methods
- `Library` system with `Book` structs (add, list, search)
- Experiment: Compare behavior with value vs pointer receivers

**Review (30 min):**
- Understand: Why does Go prefer composition? (No inheritance)

**🎯 Outcome:** Can model domains with structs; understand pointer semantics for methods.

---

### Day 6 — Interfaces & Polymorphism
**Study (90 min):**
- Implicit interfaces (Go's superpower — no `implements` keyword)
- Interface values: type + value (nil interface vs nil pointer trap)
- Type assertions (`value.(Type)`) and type switches
- Common interfaces: `Stringer`, `Reader`, `Writer`, `error`
- Empty interface (`any`) — when to use, when to avoid

**Build (120 min):**
- `Shape` interface with `Circle`, `Rectangle`, `Triangle` implementations
- `Stringer` implementation for custom types
- Polymorphic function that accepts any `Shape` and calculates total area

**Review (30 min):**
- Study the "nil interface" trap: why `var p *MyType = nil; var i interface{} = p` is not nil

**🎯 Outcome:** Understand Go's implicit interface model; can design clean abstractions.

---

### Day 7 — Week 1 Project: CLI Task Manager
**Build (240 min):**
Command-line to-do app with:
- Add, list, complete, delete tasks
- Structs: `Task` (ID, Title, Done, CreatedAt)
- Data persistence to JSON file (`encoding/json`)
- Interfaces: `Storage` (JSON file implementation, in-memory for testing)
- Error handling throughout

**Review (60 min):**
- Refactor for readability
- Add `go doc` comments
- Write a proper README
- **Commit and push to GitHub**

**🎯 Outcome:** First complete Go application with structs, interfaces, file I/O, and clean architecture.

---

## 📅 WEEK 2: Intermediate Go, Concurrency & Testing (Days 8–14)
**Theme:** *"Master the idioms that make Go, Go"*

### Day 8 — Error Handling: Go's Philosophy
**Study (90 min):**
- The `error` interface (`Error() string`)
- `fmt.Errorf` with wrapping (`%w`)
- `errors.Is()` and `errors.As()` (error inspection)
- Sentinel errors (`var ErrNotFound = errors.New(...)`)
- Custom error types (when to use structs vs simple errors)
- **Never ignore errors** — always check `if err != nil`

**Build (120 min):**
- File processor with detailed, wrapped error chains
- Custom error type `ValidationError` with field-level details
- Error handling exercises: unwrap, inspect, compare

**Review (30 min):**
- Read: [Error handling in Go](https://go.dev/blog/go1.13-errors)

**🎯 Outcome:** Write robust, inspectable error chains; understand Go's explicit error philosophy.

---

### Day 9 — Packages, Modules & Project Layout
**Study (90 min):**
- `go.mod`, `go.sum`, semantic versioning
- Exported vs unexported identifiers (capitalization rule)
- Standard Go project layout (`cmd/`, `internal/`, `pkg/`)
- Workspace mode (`go.work`) for multi-module repos
- `go mod download`, `go mod verify`

**Build (120 min):**
- Refactor Week 1 Task Manager into packages:
  - `internal/models` — structs
  - `internal/storage` — persistence
  - `internal/cli` — command handling
  - `cmd/taskmanager` — main entry point

**Review (30 min):**
- Run `go mod tidy`, `go vet`, `gofmt` on the project

**🎯 Outcome:** Understand Go module system; can organize code into maintainable packages.

---

### Day 10 — Goroutines: Go's Superpower
**Study (90 min):**
- Goroutines: lightweight threads managed by Go runtime
- The `go` keyword: `go functionName()`
- Go scheduler: M:N threading model (conceptual understanding)
- `sync.WaitGroup` for coordinating goroutines
- Goroutine leaks: what they are, how to avoid them

**Build (120 min):**
- Launch 1000 goroutines that print their ID
- Concurrent URL fetcher (fetch 10 URLs simultaneously)
- Download simulator with `sync.WaitGroup`

**Review (30 min):**
- Run `go run -race` on your concurrent code
- Understand: Goroutines vs OS threads (memory footprint, scheduling)

**🎯 Outcome:** Can write concurrent programs; understand goroutine lifecycle and coordination.

---

### Day 11 — Channels & Select
**Study (90 min):**
- Channels: `make(chan Type)`, send/receive, closing
- Buffered vs unbuffered channels (synchronous vs asynchronous)
- `select` statement: multiplexing channel operations
- Channel direction: `chan<-` (send-only), `<-chan` (receive-only)
- Common patterns: worker pools, fan-in, fan-out

**Build (120 min):**
- Producer-consumer pattern with channels
- Worker pool that processes jobs from a channel
- Fan-out/fan-in: split work across workers, merge results

**Review (30 min):**
- Experiment: What happens when you send to a closed channel? (panic)
- Experiment: What happens with `select` and multiple ready channels?

**🎯 Outcome:** Can design concurrent systems using channels; understand channel semantics deeply.

---

### Day 12 — Synchronization, Context & Race Detection
**Study (90 min):**
- `sync.Mutex` and `sync.RWMutex` (when to use which)
- Race conditions: what they are, how to detect them
- `go run -race` and `go test -race`
- `context` package: cancellation, timeouts, deadlines
- `context.WithCancel`, `context.WithTimeout`, `context.WithDeadline`
- Propagating cancellation through call chains

**Build (120 min):**
- Fix a deliberately racy counter program (add mutex)
- Add context-based cancellation to a long-running task
- Build a concurrent cache with `sync.RWMutex`

**Review (30 min):**
- Run `go test -race` on all your concurrent code
- Read: [Go Concurrency Patterns: Context](https://go.dev/blog/context)

**🎯 Outcome:** Can write safe concurrent code; understand context as Go's cancellation mechanism.

---

### Day 13 — Testing: Go's Culture
**Study (90 min):**
- `testing` package: `TestXxx`, `BenchmarkXxx`, `ExampleXxx`
- Table-driven tests (Go's idiomatic pattern)
- Subtests (`t.Run`) for organization
- Test coverage: `go test -cover`, `go test -coverprofile`
- Mocking via interfaces (no mocking frameworks needed)
- `testify` library (optional but popular)

**Build (120 min):**
- Write table-driven tests for Task Manager functions
- Mock the `Storage` interface for isolated unit tests
- Benchmark slice `append` vs pre-allocation
- Aim for >80% coverage on core logic

**Review (30 min):**
- Run `go test -race` to catch data races in tests
- Review: Are your tests readable? (Good test names, clear structure)

**🎯 Outcome:** Write idiomatic table-driven tests; understand Go's testing-first culture.

---

### Day 14 — Week 2 Project: Concurrent File Processor
**Build (240 min):**
CLI tool that:
- Scans a directory recursively
- Processes files concurrently using a worker pool
- Computes checksums (SHA-256) or word counts
- Uses `sync.WaitGroup` + channels for coordination
- Supports context cancellation (Ctrl+C graceful shutdown)
- Has comprehensive unit tests (>80% coverage)
- Uses `go run -race` to verify safety

**Review (60 min):**
- Add README with architecture explanation
- Document concurrency design decisions
- **Commit and push to GitHub**

**🎯 Outcome:** Production-quality concurrent CLI tool with tests and race-free guarantees.

---

## 📅 WEEK 3: Web Development, APIs & Databases (Days 15–21)
**Theme:** *"Build real backend services"*

### Day 15 — net/http & JSON Fundamentals
**Study (90 min):**
- `net/http` package: `http.Handler`, `http.HandlerFunc`
- `http.ServeMux` (Go 1.22+ method-based routing)
- Request/response lifecycle
- JSON encoding/decoding: `encoding/json`, struct tags
- `json.Marshal`, `json.Unmarshal`, `json.NewEncoder`

**Build (120 min):**
- Basic HTTP server with JSON responses
- CRUD endpoints for a simple resource (e.g., `Book`)
- Custom JSON marshaling for time fields
- Graceful shutdown with `http.Server` + `context`

**Review (30 min):**
- Test endpoints with `curl` or Postman
- Compare `ServeMux` vs router libraries

**🎯 Outcome:** Can build basic HTTP servers with JSON APIs; understand request lifecycle.

---

### Day 16 — Routing, Middleware & Chi Router
**Study (90 min):**
- Why use a router library? (`chi`, `gorilla/mux`, `echo`)
- Middleware pattern: `func(http.Handler) http.Handler`
- Common middleware: logging, recovery, CORS, auth
- Chaining middleware
- Request context: `r.Context()`

**Build (120 min):**
- Set up `chi` router with RESTful routes
- Build middleware: request logging (method, path, duration)
- Build middleware: panic recovery
- Add CORS support

**Review (30 min):**
- Read `chi` source code for middleware implementation
- Compare with standard library `ServeMux`

**🎯 Outcome:** Can build production HTTP servers with routing and middleware stack.

---

### Day 17 — Databases: SQL with Go
**Study (90 min):**
- `database/sql` package: `DB`, `Tx`, `Stmt`, `Rows`
- PostgreSQL driver: `pgx` (preferred) or `lib/pq`
- Connection pooling (configured via `sql.DB`)
- Prepared statements
- Transactions: `Begin`, `Commit`, `Rollback`
- `sql.NullString`, `sql.NullInt64` for nullable fields

**Build (120 min):**
- Connect to PostgreSQL (local or Docker)
- CRUD operations for `books` table
- Transaction example: transfer between accounts
- Repository pattern: abstract DB operations

**Review (30 min):**
- Understand connection pool settings (`SetMaxOpenConns`, `SetMaxIdleConns`)
- Compare `pgx` vs `lib/pq`

**🎯 Outcome:** Can write SQL-backed Go applications with proper connection management.

---

### Day 18 — ORMs, Query Builders & Migrations
**Study (90 min):**
- GORM: models, associations, hooks, migrations
- `sqlx`: lightweight extension over `database/sql`
- Schema migrations with `golang-migrate`
- ORM vs raw SQL: when to use which
- Repository pattern with GORM

**Build (120 min):**
- Refactor raw SQL into GORM models
- Define relationships (belongs to, has many)
- Add migration for schema changes
- Build repository layer with interface abstraction

**Review (30 min):**
- Benchmark: GORM vs raw SQL for simple queries
- Understand N+1 query problem

**🎯 Outcome:** Can use GORM for rapid development; know when to drop down to SQL.

---

### Day 19 — REST API Design, Validation & Config
**Study (90 min):**
- REST conventions: resources, HTTP methods, status codes
- Request validation (manual, `go-playground/validator`)
- Structured error responses (consistent JSON format)
- Pagination: offset-based vs cursor-based
- Configuration: `os.Getenv`, `.env` files, `viper` (optional)

**Build (120 min):**
- Add validation to your API endpoints
- Implement pagination on list endpoints
- Structured error responses with error codes
- Load config from environment variables

**Review (30 min):**
- Test validation edge cases
- Review: Are your error messages helpful?

**🎯 Outcome:** Can design clean REST APIs with validation and proper error handling.

---

### Day 20 — Authentication & Security
**Study (90 min):**
- Password hashing with `bcrypt`
- JWT tokens: creation, validation, refresh
- Auth middleware: extracting and validating JWT
- Basic security practices: input sanitization, HTTPS
- Rate limiting concepts (token bucket, leaky bucket)

**Build (120 min):**
- User registration endpoint (hash password with bcrypt)
- Login endpoint (issue JWT)
- Protected routes with auth middleware
- Role-based access control (RBAC) basics

**Review (30 min):**
- Test auth flow end-to-end
- Understand JWT security best practices

**🎯 Outcome:** Can secure APIs with JWT-based authentication and middleware.

---

### Day 21 — Week 3 Project: REST API Service
**Build (240 min):**
Full CRUD REST API (e.g., "Notes" or "Bookstore") with:
- `chi` router with middleware stack (logging, recovery, CORS, auth)
- PostgreSQL + GORM for persistence
- JWT authentication (register, login, protected routes)
- Input validation on all endpoints
- Structured error responses
- Pagination on list endpoints
- Unit tests for handlers and services (>80% coverage)
- Environment-based configuration

**Review (60 min):**
- Add OpenAPI/Swagger documentation (or Postman collection)
- Write comprehensive README
- **Commit and push to GitHub**

**🎯 Outcome:** Production-ready REST API — your strongest portfolio piece.

---

## 📅 WEEK 4: Production Readiness, Modern Go & Job Prep (Days 22–30)
**Theme:** *"Think like an engineer from day one"*

### Day 22 — Clean Architecture & Project Structure
**Study (90 min):**
- Standard Go project layout (`cmd/`, `internal/`, `pkg/`)
- Clean architecture: handler → service → repository
- Dependency injection (manual, via constructors)
- Interface-driven design for testability
- Options pattern for configurable constructors

**Build (120 min):**
- Refactor Week 3 API into clean layers:
  - `internal/handler` — HTTP handlers
  - `internal/service` — business logic
  - `internal/repository` — data access
  - `internal/domain` — models and interfaces

**Review (30 min):**
- Verify: Can you swap the repository implementation without touching handlers?

**🎯 Outcome:** Can design maintainable, testable Go applications with clear separation of concerns.

---

### Day 23 — Structured Logging with slog
**Study (90 min):**
- `log/slog` (Go 1.21+): structured logging in standard library
- `TextHandler` vs `JSONHandler`
- Typed attributes: `slog.String`, `slog.Int`, `slog.Any`
- Request-scoped logging with `context.Context`
- Custom handlers and `LogValuer` interface
- Why `slog` over `logrus`/`zap` for new projects

**Build (120 min):**
- Add `slog` JSON logging to your API
- Request logging middleware with method, path, duration, status
- Error logging with structured context
- Configure log level from environment

**Review (30 min):**
- Compare output: text vs JSON handlers
- Understand: Why structured logging matters for observability

**🎯 Outcome:** Can implement production-grade structured logging with `slog`.

---

### Day 24 — Docker & Deployment
**Study (90 min):**
- Multi-stage Dockerfile for Go (builder + distroless/scratch)
- `docker-compose` for app + PostgreSQL
- Environment-based config for different environments
- Health checks (`/health` endpoint)
- Graceful shutdown handling

**Build (120 min):**
- Write optimized `Dockerfile` for your API
- Create `docker-compose.yml` (app + PostgreSQL)
- Add `/health` and `/ready` endpoints
- Test: `docker-compose up` and verify everything works

**Review (30 min):**
- Compare image sizes: single-stage vs multi-stage
- Understand: Why `CGO_ENABLED=0` and static binaries matter

**🎯 Outcome:** Can containerize Go applications and run them locally with Docker Compose.

---

### Day 25 — Advanced Concurrency Patterns
**Study (90 min):**
- Pipeline pattern: stages connected by channels
- Fan-out / fan-in: distributing work, collecting results
- Rate limiting with `time.Ticker` and channels
- `errgroup` (golang.org/x/sync) for concurrent error handling
- Goroutine leak detection
- Common interview patterns

**Build (120 min):**
- Implement pipeline: read → process → write
- Implement fan-out/fan-in with error handling
- Build a rate-limited API client

**Review (30 min):**
- Run `go test -race` on all concurrent code
- Review: Where could goroutines leak?

**🎯 Outcome:** Can implement advanced concurrency patterns; prepared for concurrency interview questions.

---

### Day 26 — Generics & Modern Go
**Study (90 min):**
- Go generics (Go 1.18+): type parameters, constraints
- `constraints.Ordered`, `comparable`
- When to use generics vs interfaces
- `slices` and `maps` packages (Go 1.21+)
- Generic data structures: stack, queue, set

**Build (120 min):**
- Generic `Filter`, `Map`, `Reduce` functions
- Generic stack or queue implementation
- Refactor a function to use generics where appropriate

**Review (30 min):**
- Understand: "Start simple — don't overuse generics"
- Compare: Generic version vs interface version for your use case

**🎯 Outcome:** Can use generics appropriately; understand when they add value vs complexity.

---

### Day 27 — Integration Testing & CI/CD
**Study (90 min):**
- Integration testing with test database
- `testcontainers-go` for isolated DB tests
- GitHub Actions: `go test`, `go vet`, `golangci-lint`
- Code coverage reporting
- `go fix` and `go vet` for code quality

**Build (120 min):**
- Add integration tests for your API endpoints
- Set up GitHub Actions workflow:
  - Run tests with race detector
  - Run `golangci-lint`
  - Build Docker image
- Add coverage badge to README

**Review (30 min):**
- Verify CI passes on push
- Review: Are integration tests isolated and deterministic?

**🎯 Outcome:** Can set up CI/CD pipelines; write integration tests for real APIs.

---

### Day 28 — gRPC & Protocol Buffers
**Study (90 min):**
- Protocol Buffers: `.proto` files, message types
- gRPC services: unary, streaming
- Generating Go code from `.proto`
- gRPC interceptors (middleware equivalent)
- When to use gRPC vs REST

**Build (120 min):**
- Define a `.proto` service (e.g., `UserService`)
- Generate Go code with `protoc`
- Implement gRPC server and client
- Add basic interceptor for logging

**Review (30 min):**
- Compare: gRPC vs REST for your use case
- Understand: HTTP/2, binary protocol benefits

**🎯 Outcome:** Can build basic gRPC services; understand when gRPC is the right choice.

---

### Day 29 — Capstone Project: Build Day
**Build (240 min):**
Start or continue a capstone that combines everything:
- **URL Shortener with Analytics** (recommended):
  - REST API for creating/redirecting short URLs
  - PostgreSQL for persistence
  - JWT auth
  - Background worker (goroutine) for click analytics
  - Docker + docker-compose
  - `slog` structured logging
  - Comprehensive tests
  - CI/CD with GitHub Actions

**Review (60 min):**
- Architecture review: draw component diagram
- Ensure clean architecture layers

**🎯 Outcome:** A capstone project that demonstrates all learned skills.

---

### Day 30 — Polish, Portfolio & Interview Prep
**Morning (120 min):**
- Finish capstone project
- Add comprehensive README with:
  - Architecture diagram
  - API documentation
  - Setup instructions
  - Testing instructions
- Clean commit history (`git rebase -i` if needed)
- Push to GitHub

**Afternoon (120 min):**
- Review common Go interview questions:
  - Slices vs arrays (backing arrays, capacity)
  - Goroutine leaks and how to prevent them
  - Channel deadlocks
  - Interface internals (type + value)
  - Garbage collection basics
  - Context cancellation patterns
- Solve 5 LeetCode medium problems in Go
- Mock interview: explain your capstone project aloud
- Update LinkedIn with Go skills

**🎯 Outcome:** Polished GitHub portfolio, interview-ready explanations, clear next steps.

---

## 📊 Weekly Milestones & Deliverables

| Week | Focus | Tangible Deliverable | Skills Demonstrated |
|------|-------|---------------------|---------------------|
| **Week 1** | Language Foundations | CLI Task Manager (GitHub) | Core syntax, structs, interfaces, file I/O |
| **Week 2** | Concurrency & Testing | Concurrent File Processor (GitHub) | Goroutines, channels, sync, context, tests |
| **Week 3** | Web & Databases | REST API Service (GitHub) | HTTP, routing, PostgreSQL, auth, validation |
| **Week 4** | Production & Modern Go | Capstone + CI/CD (GitHub) | Docker, slog, gRPC, generics, clean architecture |

---

## 📚 Recommended Resources

| Type | Resource | Purpose |
|------|----------|---------|
| **Official** | [A Tour of Go](https://go.dev/tour/) | Interactive basics |
| **Official** | [Effective Go](https://go.dev/doc/effective_go) | Idiomatic patterns |
| **Official** | [Go by Example](https://gobyexample.com/) | Quick syntax reference |
| **Official** | [Go Standard Library Docs](https://pkg.go.dev/std) | API reference |
| **Book** | *"The Go Programming Language"* (Donovan & Kernighan) | Deep understanding |
| **Book** | *"Learning Go"* (Jon Bodner) | Modern Go (1.18+) |
| **Book** | *"Concurrency in Go"* (Katherine Cox-Buday) | Goroutines mastery |
| **Practice** | [Exercism Go Track](https://exercism.org/tracks/go) | Daily exercises |
| **Practice** | [Gophercises](https://gophercises.com/) | Project-based learning |
| **Community** | r/golang, Gopher Slack | Code review feedback |

---

## ⚠️ Critical Pitfalls to Avoid

1. **Don't write Java/Python in Go** — Embrace simplicity. No classes, no inheritance, no generics overuse.
2. **Never ignore errors** — Always check `if err != nil`. It's idiomatic, not optional.
3. **Don't fear concurrency** — Start with goroutines, add channels when needed. Use `go test -race`.
4. **Avoid premature optimization** — Write clear code first, benchmark later.
5. **Don't skip testing** — Table-driven tests are a Go superpower. Use them from Day 1.
6. **Read others' code** — Study Docker, Kubernetes, Caddy source for real-world patterns.
7. **Commit daily** — Your GitHub commit history becomes your portfolio.

---

## ✅ Job-Ready Checklist (By Day 30)

- [ ] Explain Go's memory model and garbage collection
- [ ] Build concurrent programs with goroutines and channels safely
- [ ] Design clean REST APIs with proper HTTP semantics
- [ ] Write testable code with >80% coverage (unit + integration)
- [ ] Containerize applications with Docker (multi-stage builds)
- [ ] Implement structured logging with `slog`
- [ ] Debug race conditions using `go run -race`
- [ ] Use `go vet`, `gofmt`, `golangci-lint` for code quality
- [ ] Explain when to use pointers vs values, interfaces vs concrete types
- [ ] Deploy a complete system locally with Docker Compose
- [ ] 4+ projects on GitHub with clean READMEs and commit history
- [ ] Can whiteboard common concurrency patterns

---

## 🎯 Why This Unified Roadmap Works

1. **Correct Dependency Chain:** Concurrency (Week 2) → HTTP (Week 3). You cannot understand `net/http` without understanding goroutines first.
2. **Early Pointers:** Day 5 introduces pointers before interfaces, so interface value vs pointer semantics make sense.
3. **Testing from Day 13:** Not an afterthought — integrated into the concurrent file processor project.
4. **Modern Go:** `slog` (Go 1.21+), generics (Go 1.18+), Go 1.22 `ServeMux` improvements.
5. **Realistic Scope:** No microservices fantasy in 30 days. One solid monolith beats three broken services.
6. **Portfolio-Focused:** Every week ends with a GitHub-ready project.

> **"The best time to start was yesterday. The second best time is now."** — Start Day 1 today.
