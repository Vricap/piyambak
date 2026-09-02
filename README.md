# Piyambak

> **Self-hosted. Deployless. Anonymous. Real-time.**

Piyambak (Javanese _krama_ for **"self"**) is a lightweight real-time chat
application built with Go. Instead of deploying to a cloud server, Piyambak runs
entirely on your local machine and becomes publicly accessible through an
**Ngrok tunnel**. Share one URL, and anyone can instantly join your chat rooms.

No accounts. No email. No registration. Just enter a username and start
chatting.

---

## ✨ Features

- **Real-time messaging** powered by WebSockets.
- **Deployless public access** using the Ngrok SDK.
- Create unlimited chat rooms.
- Password-protected private chat rooms.
- Anonymous usernames — no user accounts or authentication.
- Persistent chat history stored locally with SQLite.
- Join, leave, and switch between chat rooms anytime.
- Lightweight single-binary Go application.

---

## Preview

> Example flow:

```text
Open Piyambak URL
        │
        ▼
Enter a username
        │
        ▼
Join an existing room
or create a new room
        │
        ▼
Start chatting in real time
```

---

## Tech Stack

| Technology          | Purpose                                                    |
| ------------------- | ---------------------------------------------------------- |
| **Go**              | Backend server and application logic                       |
| **Fiber**           | HTTP web framework                                         |
| **Fiber WebSocket** | Real-time bidirectional communication                      |
| **SQLite**          | Store chat rooms and chat history locally                  |
| **Ngrok SDK**       | Expose the local server to the internet without deployment |

---

## How It Works

Piyambak follows a simple self-hosting architecture.

```text
       Public Internet
              │
              ▼
     https://xxxx.ngrok.app
              │
        Ngrok Tunnel
              │
              ▼
Go + Fiber WebSocket Server
              │
         SQLite Database
```

The host runs Piyambak locally, while Ngrok securely forwards public traffic to
the local server.

**There is no centralized server.** Every instance of Piyambak is its own
independent chat server.

---

## Project Structure

```text
piyambak/
├── cmd/                # Application entrypoint
├── internal/
│   ├── chat/           # Chat room & message logic
│   ├── websocket/      # WebSocket handlers
│   ├── database/       # SQLite connection
│   └── ngrok/          # Ngrok tunnel initialization
├── views/              # HTML templates
├── public/             # Static assets
├── chat.db             # SQLite database (generated)
├── go.mod
└── main.go
```

---

## Running Piyambak

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/piyambak.git
cd piyambak
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Configure Ngrok

Create an Ngrok account and obtain your authentication token.

Set it as an environment variable through .env file.

#### .env

```bash
NGROK_AUTHTOKEN=your_ngrok_token
```

### 4. Run the server

```bash
go run .
```

You'll see something similar to:

```text
Server started on :3000

Public URL:
https://abcd-1234.ngrok.app
```

Open the generated URL in your browser and start chatting.

---

## Usage

### Join a Chat Room

1. Open the public URL.
2. Enter any username.
3. Select an existing room.
4. Start chatting.

### Create a Chat Room

1. Click **Create Room**.
2. Enter a room name.
3. (Optional) Set a password.
4. Share the room with others.

### Leave a Room

Leave at any time and join another room without creating a new identity.

---

## Data Storage

Piyambak intentionally stores **very little information**.

| Stored                     | Not Stored      |
| -------------------------- | --------------- |
| Chat room metadata         | User accounts   |
| Password for private rooms | User IP Address |
| Chat history               | User profiles   |
| Timestamps                 |                 |

All data is stored locally inside a SQLite database owned by the host.

---

## Why "Deployless"?

Traditional web applications require deploying to a VPS, cloud platform, or
container service before they become publicly accessible.

Piyambak removes that step.

```text
Run locally
     │
Ngrok SDK creates a tunnel
     │
Receive a public HTTPS URL
     │
Share the URL
     │
People join your chat room
```

The application is still self-hosted. Your computer is the server, but no
deployment infrastructure is required.

---

## Philosophy

Piyambak is built around a few simple ideas:

- **Self-hosted** — you own the server.
- **Deployless** — no VPS or cloud deployment required.
- **Anonymous** — no accounts or registration.
- **Lightweight** — minimal dependencies.

---

## Future Improvements

- [ ] End-to-end encrypted private rooms.
- [ ] Typing indicators.
- [ ] Online user list.
- [ ] Message reactions.
- [ ] File and image sharing.
- [ ] Ephemeral rooms with automatic expiration.

---

## License

MIT License.

---

## Author

Built by **Vricap**.
