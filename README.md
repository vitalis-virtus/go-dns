# Simple DNS Server in Go using miekg/dns

This project demonstrates how to build a basic DNS server using [github.com/miekg/dns](https://github.com/miekg/dns).  
It uses **in-memory storage** and provides a working example of handling DNS queries using static responses.

---

## 🚀 Getting Started

The DNS server listens on **port 9090 (UDP)** and responds to basic queries like A-record lookups.

---

## 🐳 Docker Setup

The project includes a **Dockerfile** and **docker-compose.yml** for easy deployment.  
This ensures consistent setup and isolates the DNS service from your local environment.

To run the server using Docker Compose:

1. Make sure Docker and Docker Compose are installed.
2. In the project directory, run:

   `docker-compose up --build`

This will start the DNS server on `udp://localhost:9090`.

---

## 🧪 How to Test

Once the server is running, you can test it in one of the following ways:

- Using `dig` (on Unix systems):  
  `dig @127.0.0.1 -p 9090 example.com`

- Using PowerShell (on Windows):  
  `Resolve-DnsName -Server 127.0.0.1 -Port 9090 example.com`

---

## 🔧 Internal Behavior

- All DNS records are served from a static in-memory handler.
- The response for A-type queries (e.g. `example.com`) returns a predefined IP address (like `1.2.3.4`).
- You can customize the domain handling logic in the Go code.

---
