const express = require("express");
const app = express();
const port = 6677;

app.use(express.json());

// GET route
app.get("/api/get", (req, res) => {
  res.json({
    status: "success",
    method: "GET",
    message: "GET request received",
    query: req.query,
  });
});

// POST route (returns receive time)
app.post("/api/post", (req, res) => {
  const receivedAt = new Date().toISOString(); // exact server time

  res.json({
    status: "success",
    method: "POST",
    message: "POST request received",
    receivedAt,          // returning timestamp here
    body: req.body,      // whatever client sent
  });
});

// PUT route
app.put("/api/put", (req, res) => {
  res.json({
    status: "success",
    method: "PUT",
    message: "PUT request received",
    body: req.body,
  });
});

// PATCH route
app.patch("/api/patch", (req, res) => {
  res.json({
    status: "success",
    method: "PATCH",
    message: "PATCH request received",
    body: req.body,
  });
});

// DELETE route
app.delete("/api/delete", (req, res) => {
  res.json({
    status: "success",
    method: "DELETE",
    message: "DELETE request received",
  });
});

// Default route
app.get("/", (req, res) => {
  res.send("Simple Express Backend Running Successfully 🚀");
});

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
