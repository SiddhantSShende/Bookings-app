# Ecommerce Website for Business 
This project is a comprehensive ecommerce website designed to facilitate online shopping for businesses. It includes features such as product listings, shopping cart functionality, user authentication, and payment processing.
- Uses [Chi](https://github.com/go-chi/chi/v5) router for handling HTTP requests
- Uses [nosurf](https://github.com/justinas/nosurf) for CSRF protection
- Uses [SCS](https://github.com/alexedwards/scs/v2) for session management
- Implements a modular architecture for scalability and maintainability  

## Features
- User registration and login
- Product catalog with categories and search functionality
- Shopping cart and checkout process
- Payment gateway integration
- Order management for users and admins
- Responsive design for mobile and desktop
## Technologies Used
- Frontend: HTML, CSS, JavaScript, Bootstrap
- Backend: Golang
- Database: MongoDB, PostgreSQL
- Payment Gateway: Stripe/PayPal
## Installation
  1. Clone the repository: `git clone
2. Navigate to the project directory: `cd ecommerce-website`
  3. Install dependencies: `go mod tidy`
  4. Set up the database and configure the connection settings in the config file.
  5. Run the application: `go run main.go`
6. Open your browser and go to `http://localhost:8080`
7. ## Contributing
  Contributions are welcome! Please fork the repository and create a pull request with your changes.
