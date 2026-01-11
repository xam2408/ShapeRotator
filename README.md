# 🧊 Go 4D Tesseract Engine

A lightweight 4D rendering engine built from scratch using **Go** (Backend) and **JavaScript + Canvas** (Frontend). This project demonstrates the mathematics behind visualizing four-dimensional objects on a two-dimensional screen using double perspective projection.

## 🚀 About the Project

This application renders a rotating **Tesseract (Hypercube)** by performing all mathematical calculations on the server side.

1.  **Backend (Go):** Generates vertices in 4D space, calculates rotations (including the "inside-out" rotation in the 4th dimension), projects points to 3D space, and finally projects them onto a 2D plane.
2.  **Communication:** The server exposes a REST API endpoint that streams the calculated coordinate data in JSON format.
3.  **Frontend (JS + Canvas):** Fetches the point data and renders the animation frame using the HTML5 `<canvas>` API, utilizing a pre-defined edge map.

## 🛠️ Tech Stack

* **Backend:** Go (Golang) - Standard library only (`net/http`, `math`, `encoding/json`).
* **Frontend:** HTML5 Canvas API, Vanilla JavaScript.
* **Rendering:** Real-time vector drawing (60 FPS).
* **Data Format:** JSON.

## 🧮 How It Works (The Math)

The project implements a custom graphics pipeline adapted for four dimensions:

1.  **Definition:** The Tesseract is defined as an array of 16 vertices with coordinates $(x, y, z, w)$, where each coordinate is $\pm 1$.
2.  **4D Rotation:** We apply rotation matrices in planes involving the 4th dimension (e.g., the **ZW** plane). This creates the "inside-out" transformation effect:
    * $z' = z \cdot \cos(\theta) - w \cdot \sin(\theta)$
    * $w' = z \cdot \sin(\theta) + w \cdot \cos(\theta)$
3.  **Projection 4D → 3D:** To descend one dimension, we use perspective projection by dividing by the $W$ coordinate:
    * $P_{3D} = P_{4D} \cdot \frac{1}{distance - w}$
4.  **Rotation & Projection 3D → 2D:** The resulting 3D object is rotated (X/Y/Z axes) and projected onto the screen by dividing by $Z$:
    * $x_{screen} = x \cdot \frac{1}{distance - z}$
    * $y_{screen} = y \cdot \frac{1}{distance - z}$

## 🚧 Roadmap / Todo

Current development goals and features in progress:

- [ ] **Mouse Interaction:** Implement logic to rotate the Tesseract by dragging the mouse (sending offset data to the Go backend).
- [ ] **Depth Cues (Z/W-Depth):** Adjust the brightness or opacity of points based on their depth (Z or W coordinates) to enhance the 3D/4D perception.
- [ ] **More 4D Shapes:** Implementation of the Pentachoron (5-cell) or the 24-cell.
- [ ] **Optimization:** Move the static edge definitions back to the backend (Go) to support dynamic topology changes (e.g., changing shapes on the fly).

## 📂 Project Structure

```text
├── main.go       # Server logic, 4D math calculations, HTTP handling
├── index.html    # Visual layer (Canvas) and edge definition (JS)
└── README.md     # Project documentation
