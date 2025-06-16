
# Mitti & More – Golang E-commerce Backend

<table>
  <tr>
    <td>
      <img src="https://upload.wikimedia.org/wikipedia/commons/4/4e/Docker_%28container_engine%29_logo.svg" alt="Docker Logo" width="64" />
    </td>
    <td style="vertical-align: middle; padding-left: 10px;">
      <a href="https://github.com/sumedhahire/golang-Shop-Backend/pkgs/container/golang-shop-backend" target="_blank" style="font-size: 18px; font-weight: bold; text-decoration: none; color: #0366d6;">
        GoLang Shop Backend Docker Container
      </a>
    </td>
  </tr>
</table>

A production-ready backend for an online store, built using Go. Features include custom OAuth2.0 authentication, Razorpay integration, email notifications, and modular API structure. Built with scalability and maintainability in mind.

---

## 📦 Tech Stack

- **Language:** Go 1.21+
- **Framework:** Echo
- **Database:** MySql
- **ORM:** Entgo
- **Migrations:** Goose
- **Payments:** Razorpay
- **Auth:** Custom OAuth2.0 (access & refresh tokens)
- **Containerization:** Docker

---

## ✅ Features

- OAuth2.0-based authentication system (token generation, refresh, expiry handling)
- Razorpay integration with secure signature verification
- RESTful APIs for product, tag, order, and user management
- Role-based access logic 
- Clean separation of handler, service, and model layers

---

## 🚀 Quickstart

### 1. Clone & Configure

```bash
git clone https://github.com/sumedhahire/golang-Shop-Backend.git
cd golang-Shop-Backend
cp .env.example .env