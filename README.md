# Bootstrap Admin Template with Golang Fiber

In this project, we explore the relatively uncommon integration of Golang Fiber with a Bootstrap-based admin template. By utilizing the free Bootstrap template from [Star Admin Free](https://www.bootstrapdash.com/product/star-admin-free), we aim to implement and seamlessly integrate its front-end design with Golang Fiber for a streamlined and modern web development experience.

Please note that this project focuses solely on the integration of Fiber with HTML and templating using an HTML engine. It does not cover any aspects related to CRUD operations or database interactions.


## Getting Started

### 1. Configure Environment Variables
Create a `.env` file in the root directory of the project and fill it with your configuration settings with basic values from `.example.env`.

### 2. Install Dependencies
Run the following command to ensure all necessary modules are installed:

```bash
go mod tidy
```

### 3. Start the Development Server
To start the development server, run:

```bash
go run main.go
```

This will start the server and automatically load changes when you rerun the command after making changes.

Alternatively, if you prefer using air for live reloading during development, simply run:

```bash
air
```

Make sure to configure air according to your project's needs by adjusting the settings in the .air.toml file.

### 4. Start the Production Server
To start the server in production mode, you can build the binary and run it:

#### On Windows:
```bash
go build -o lorem.exe
lorem.exe
```

#### On Linux/macOS:
```bash
go build -o lorem
./lorem
```
