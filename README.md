# OGPer - OGP Card Generator

OGPer is an open-source tool for easily generating and managing OGP (Open Graph Protocol) cards.
https://ogper.pages.dev


## Key Features

- Simple UI for OGP card creation
- Instant preview of generated cards
- Responsive design
- Supports image uploads up to 1MB
- Easy sharing of generated cards

## Usage

https://github.com/user-attachments/assets/44ef8222-9834-470d-aea2-a6f6f0446ec2
below is generated page with OGP in this video
https://ogper.onrender.com/pd3ae4d

## Directory Structure

```
ogper/
├── backend/            # Backend related files
│   ├── config/         # Configuration files
│   ├── domain/         # Domain models
│   ├── handler/        # HTTP handlers
│   ├── middleware/     # Middleware
│   ├── repository/     # Data access layer
│   ├── usecase/        # Business logic
│   └── main.go         # Entry point
├── frontend/           # Frontend related files
│   ├── public/         # Static files
│   ├── src/            # Source code
│   │   ├── assets/     # Asset files
│   │   ├── components/ # React components
│   │   ├── pages/      # Page components
│   │   └── App.tsx     # Main application
│   └── vite.config.ts  # Vite configuration
└── README.md
```

## Setup Instructions

1. Clone the repository
   ```bash
   git clone https://github.com/your-username/ogper.git
   cd ogper
   ```

2. Backend setup
   ```bash
   cd backend
   go mod download
   ```

3. Frontend setup
   ```bash
   cd ../frontend
   npm install
   ```

4. Environment variables configuration
   ```bash
   cp .env.example .env
   # Edit .env file
   ```

5. Start the application
   ```bash
   # Backend
   cd ../backend
   go run main.go

   # Frontend (in another terminal)
   cd ../frontend
   npm run dev
   ```
