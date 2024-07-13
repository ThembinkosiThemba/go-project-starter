# 🚀 Golang Project Starter Kit

This is the Golang Project Starter Kit! 🎉 This a simple repository that contains starter code for a Golang project, following a Domain-Driven Design (DDD) approach. It simplifies the process of setting up a new project from scratch, allowing you to focus on building your application's core functionality.

In this setup, we use the `User` entity as an example, but you can easily replicate and add new entities as required for your specific project needs.

## 🌟 Features

- 🏗️ Domain-Driven Design architecture
- ✅ Complete CRUD (Create, Read, Update, Delete) operations
- 📅 Event tracking using Mixpanel
- 🗄️ Database setup and integration
  - MongoDB support
  - PostgreSQL support with advanced features:
    - Transaction handling
    - Prepared statements
    - Database migrations
- 🌐 HTTP REST APIs using Gin-Gonic framework
  - Custom response handling
- 🛡️ Basic input validation
- 🧩 Modular and extensible codebase

## 🚀 Getting Started

1. Clone the repository:

```bash
git clone https://github.com/ThembinkosiThemba/go-project-starter.git
cd golang-project-starter
```
2. Install dependencies and running:

```bash
go mod tidy
make run
```

3. Set up your environment variables (copy `.env.example` to `.env` and fill in your variables)

## 🏗️ Project Structure

## 💾 Database Setup

### MongoDB

1. Ensure you have MongoDB installed and running
2. Update the MongoDB connection string in your `.env` file

### PostgreSQL

1. Install PostgreSQL if you haven't already
2. Create a new database for your project
3. Update the PostgreSQL connection details in your `.env` file
4. The migrations will automatically run when you run the project.

### Mixpanel
Login in to [Mixpanel](mixpanel.com) and create a project, get the project id in the settings and update your env file as well.

## 🛠️ Customization

To add a new entity:

1. Create a new file in `internal/entity` for your entity
2. Implement repository interfaces in `internal/infrastructure` and choose either database.
3. Create use cases in `internal/application/usecase`
4. Add HTTP handlers in `internal/routes/handler`
5. Update routes in `internal/routes/handler/routes.go`

## 🧪 Testing

Coming soon...

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
