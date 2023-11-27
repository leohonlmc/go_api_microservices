# Go API Microservices

<p>
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" height="25"/>
  <img src="https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white" height="25"/>
</p>

<p>
  <img src="./go.png" alt="Go Bear Logo" width="250" align="right"/>
  <b>Project Overview</b><br>
  `go_api_microservices` is a project developed in Go, utilizing Gin for handling routing and MongoDB for storing data. It offers API endpoints that allow data retrieval from the inventories and news collections. The data provided by these endpoints consist of simulated information about inventories from local food banks in Canada and news feeds.
</p>

### Features

- RESTful API endpoints to interact with MongoDB.
- Utilizes Gin web framework for routing.
- MongoDB integration using Go driver.
- Clean and structured codebase.

## Getting Started

### Prerequisites

- Go (1.x or later)
- MongoDB
- Git

### Installation

1. **Clone the Repository**

2. **Set Up MongoDB**

- Ensure MongoDB is running.
- Use the `inventories` and `news` collections.

3. **Configure Environment Variables**
   Create a `.env` file in the root with your MongoDB URI:

4. **Run the Application**

## Usage

### API Endpoints

- **Get Inventories**: `GET /inventories`

  - Fetches documents from the `inventories` collection.
  - Supports optional query parameters `location` and `category` for filtering.

- **Get News**: `GET /news`
  - Retrieves documents from the `news` collection.

### Example Requests

- Fetching inventories with filtering:
  curl "https://go-api-microservices.onrender.com/inventories?location=Scarborough&category=Beverages"

- Fetching news:
  curl https://go-api-microservices.onrender.com/news

## Contributing

Contributions are welcome. To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b my-new-feature`).
3. Commit your changes (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin my-new-feature`).
5. Create a new Pull Request.
