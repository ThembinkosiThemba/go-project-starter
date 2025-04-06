## Golang Project Starter Kit

This a repository that contains starter code for a Golang project, following a Domain-Driven Design (DDD) approach. It simplifies the process of setting up a new project from scratch, allowing you to focus on building your application's core functionality.

In this setup, we use the `User` entity as an example, but you can easily replicate and add new entities as required for your specific project needs.

### Table of Contents

1. [Features](#-features)
2. [Getting Started](#-getting-started)
3. [Project Setup](#-project-setup)
4. [Database Setup](#-database-setup)
   - [MongoDB](#mongodb)
   - [PostgreSQL](#postgresql)
5. [Docker Setup](#docker-setup)
6. [Mixpanel](#mixpanel)
7. [Logger](#logger)
8. [Customization](#-customization)
9. [Testing](#-testing)
10. [Contributing](#-contributing)
11. [License](#-license)

### 🌟 Features

- 🏗️ Domain-Driven Design architecture
- ✅ Complete CRUD (Create, Read, Update, Delete) operations
- 📅 Event tracking using Mixpanel
- 🐳 Docker support which makes it easy to deploy project on multiple platforms like GCP, AWS etc
- 🗄️ Database setup and integration
  - 🍃 MongoDB support
  - 🐬My SQL and 🐘PostgreSQL support with advanced features:
    - Transaction handling
    - Prepared statements
    - Database migrations
- 🌐 HTTP REST APIs using Gin-Gonic framework
  - Custom response handling
- 🛡️ Basic input validation

### Getting Started

1. Clone the repository:

```bash
git clone https://github.com/ThembinkosiThemba/go-project-starter.git
cd golang-project-starter
```

2. Install dependencies and run the roject:

```bash
go mod tidy
make run
```

If you don't have `Make` installed, you can run

```bash
go run cmd/main.go
```

Or install `Make` by running

```bash
sudo apt install make
```

3. Set up your environment variables (copy `.env.example` to `.env` and fill in your variables)

### 🏗️ Project Setup

#### 💾 Database Setup

Depending on which database you are going to be using, make sure you update the `main.go` initialization lines so it works perfectly for your choice. By default, this project uses `Mongo DB` and this code is as follows:

```golang
userRepo, err := config.InitializeRepositoriesMongo()
if err != nil {
  log.Fatal(err)
}

userUsecase := config.InitializeUsecasesMongo(userRepo)
```

Notice we are using these two mongo functions which are `InitializeRepositoriesMongo` and `InitializeUsecasesMongo`

If for example you want to use Postgres, you will update these functions and use `InitializeRepositoriesPostgres` and `InitializeUsecasesPostgres`:

```golang
userRepo, err := config.InitializeRepositoriesPostgres()
if err != nil {
  log.Fatal(err)
}

userUsecase := config.InitializeUsecasesPostgres(userRepo)

```

### MongoDB

1. Ensure you have MongoDB installed and running locally (or via docker). Alternatively, you can use [Mongo DB Atlas](https://www.mongodb.com/cloud/atlas/register), create a project, and get the connection string.
2. Update the MongoDB connection string in your `.env` file

### PostgreSQL

1. Install PostgreSQL if you haven't already.
2. Create a new database for your project.
3. Update the PostgreSQL connection details in your `.env` file.
4. Use the migration endpoint via curl or postman to run all necessary migrations.

Alternatively, you can use solutions like [Aiven](https://aiven.io/) which has completely hosted db solutions.

### MYSQL

1. Install MYSQL if you haven't already. You can also run mysql via docker and use that connection instead.
2. Create a new database for your project.
3. Update the MYSQL connection details in your `.env` file.
4. Use the migration endpoint via curl or postman to run all necessary migrations.

### Docker setup

Open the [Dockerfile](Dockerfile) and rename make changes to the following line (`/go-project-starter`) to reflect the name of the project you are building.

```Dockerfile
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-project-starter ./cmd/main.go

and

CMD [ "/go-project-starter" ]
```

### Mixpanel

This project also has support for event tracking using Mixpanel. Login to [Mixpanel](mixpanel.com) and create a project, get the project id in the settings and update your env file as well.

### Emails

This project also has support for emailing using Sendgrid. Login to [Sendgrid](https://login.sendgrid.com/) and create a project, get the api key in the settings and update your env file as well.

### Logger

This projects now supports a custom logger. Features include:

- saving logs to files (errors, warnings and infos). These logs can be stored in their separate files. Check `/logs` folder once you use any of the logs.
- print's out logs in the terminal

## 🛠️ Customization

To add a new entity:

1. Create a new file in `internal/entity` for your entity
2. Implement repository interfaces in `internal/repository` and choose either database.
3. Create use cases in `internal/application/usecase`
4. Add HTTP handlers in `internal/routes/handler`
5. Update routes in `internal/routes/handler/routes.go`

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
