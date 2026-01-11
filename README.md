# 🧊 Go 3D Wireframe Objects

A lightweight 3D rendering engine built from scratch using **Go** (Backend) and **JavaScript** (Frontend). This project demonstrates how 3D math, rotation matrices, and perspective projection work without relying on external graphics libraries like Three.js or OpenGL.

## 🚀 About the Project

This application renders a rotating 3D objects by performing all mathematical calculations on the server side.
1.  **Backend (Go):** Calculates 3D point rotations (X and Y axes) and projects them onto a 2D plane.
2.  **Communication:** The server exposes a REST API endpoint that streams coordinate data in JSON format.
3.  **Frontend (JS + SVG):** Fetches the data and updates the DOM elements (lines and circles) within an `<svg>` canvas.

## 🛠️ Tech Stack

* **Backend:** Go (Golang) - Standard library only (`net/http`, `math`, `encoding/json`).
* **Frontend:** HTML5, CSS3, Vanilla JavaScript.
* **Rendering:** SVG (Scalable Vector Graphics).
* **Data Format:** JSON.

## 🧮 How It Works (The Math)

The project implements a classic 3D graphics pipeline from scratch:

1.  **Definition:** The cube is defined as an array of 8 vertices (X, Y, Z coordinates).
2.  **Rotation:** Rotation matrices are applied for the X and Y axes:
    * `y' = y*cos(θ) - z*sin(θ)`
    * `z' = y*sin(θ) + z*cos(θ)`
3.  **Projection:** To achieve the 3D effect (perspective), we use a "Weak Perspective Projection" by dividing by Z:
    * `x_screen = x / (distance - z)`
    * `y_screen = y / (distance - z)`

## 📂 Project Structure

```text
├── main.go       # Server logic, math calculations, HTTP handling
├── index.html    # Visual layer (SVG) and data fetching script (JS)
└── README.md     # Project documentation
