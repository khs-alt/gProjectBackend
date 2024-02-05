# Ghosting Artifact Measurement Service - Backend

## Overview
This repository contains the backend part of the Ghosting Artifact Measurement Service, a web service developed for a research project funded by Google Korea. The service aims to objectively measure ghosting artifacts in videos and HDR images, supporting the creation of labeled datasets for Machine Learning model training.

## Backend Features
- **RESTful API**: Developed in Go, providing endpoints for data management and interaction with the front-end.
- **Concurrent User Handling**: Supports over 50+ concurrent users for efficient data labeling and collection.
- **Database Management**: Utilizes MariaDB for robust data storage and retrieval.
- **Concurrent Processing**: Implements goroutines for handling frequent database inserts, reducing data loss and improving performance.

## Technical Stack
- **Programming Language**: Go
- **Database**: MariaDB
- **Concurrency**: Goroutines for concurrent processing
- **Deployment**: Deployed on Google Cloud Platform (GCP) for scalability and reliability.

## Getting Started

### Prerequisites
Ensure you have the following installed:
- Go (latest version recommended)
- MariaDB (10.x or newer)

### Setup
1. **Clone the repository**
    ```bash
    git clone <this-repository-url>
    cd into-the-cloned-directory
    ```

2. **Database Configuration**
    - Install and configure MariaDB on your server.
    - Create a database and user for this project.
    - Import the initial schema provided in the `database/schema.sql` file.

3. **Environment Variables**
    - Set up the necessary environment variables for database connections and any other configurations.

4. **Running the Server**
    ```bash
    go build -o ghosting-artifact-server
    ./ghosting-artifact-server
    ```

## Development Guidelines

- **Code Structure**: Follow Go best practices and project-specific guidelines for structuring your code.
- **Commit Messages**: Write clear, concise commit messages that describe the changes made.
- **Pull Requests**: For new features or bug fixes, submit pull requests for review.

## Contributing
Contributions to improve the backend are welcome. Please fork the repository and submit pull requests with your changes. For major changes or new features, please open an issue first to discuss the proposal.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements
Special thanks to Google Korea for funding this project and to all team members who have contributed to the development of this service.
