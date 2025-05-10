
# Basic Realtime Chat

Connect instantly. Chat without limits!  
**Secure, Fast, Seamless.**

Stay connected with friends, teams, or communities through our powerful chat — featuring real-time communication, modern UI, and scalable backend.

![Hero Screenshot](./demo/image_demo.png)

---

## 🛠 Tech Stack

### Frontend
- [Next.js 15](https://nextjs.org/)
- Tailwind CSS
- TypeScript

### Backend
- [Go (Golang)](https://go.dev/)
- [Fiber v2](https://docs.gofiber.io/)
- [GORM](https://gorm.io/)
- Clean Architecture

### Database
- PostgreSQL

### Realtime
- WebSocket (Real-time chat message delivery)

---

## ⚙️ Project Structure

### Frontend (Next.js 15)
- Located in the `frontend/` folder
- Handles user interface, routing, and communicates with backend via REST and WebSocket

### Backend (Go)
- Located in the `backend/` folder
- Clean architecture layers: `delivery`, `usecase`, `repository`
- RESTful API and WebSocket endpoint
- PostgreSQL with GORM for ORM

---

## 🚀 Features

- ✅ Real-time messaging with WebSocket
- ✅ Secure user authentication
- ✅ Clean and scalable architecture
- ✅ Responsive UI
- ✅ Easy deployment setup
- ✅ PostgreSQL integration
- ✅ Production-ready structure

---

## 📦 Installation

### Prerequisites
- Go 1.20+
- Node.js 20+ or [Bun](https://bun.sh/)
- PostgreSQL
- Docker (optional for local development)

### Backend Setup

```bash
cd backend
go run main.go
````

### Frontend Setup

```bash
cd frontend
npm install # or bun install
npm run dev # or bun dev
```

---

## 🔄 WebSocket Usage

* WebSocket endpoint: `ws://localhost:PORT/api/v1/rooms/:room_code/ws`
* Events supported:
  * `new_message`: receive and broadcast messages in real time
  * `member_online`: notify when a member joins the room
  * `member_offline`: notify when a member leaves the room


## 📁 Example `.env` for Backend

```env
FIBER_HOST=0.0.0.0
FIBER_PORT=8000

DB_HOST=0.0.0.0
DB_PORT=5432
DB_DATABASE=test
DB_USERNAME=test
DB_PASSWORD=123456789
DB_SSL_MODE=disable

JWT_SECRET=secret
JWT_EXPIRE=1 # hours
JWT_REFRESH_SECRET=refresh_secret
JWT_REFRESH_EXPIRE=24 # hours
```

## 📸 Demo Video

![Demo Video](./demo/video_demo.mp4)

## 🧪 Coming Soon

* User profile and avatars
* File/image sharing in chat
* Message history and pagination
* Admin dashboard

## 🤝 Contributing

We welcome contributions! Please fork the repo, create a feature branch, and open a pull request.