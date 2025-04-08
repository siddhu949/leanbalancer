const express = require('express');
const cors = require('cors');

const app = express();

// Allow requests from your frontend
app.use(cors({ origin: 'http://localhost:3000' }));

// If you need to allow multiple origins
// app.use(cors({ origin: ['http://localhost:3000', 'http://localhost:9002'] }));

// Start the server
app.listen(9002, () => {
  console.log('Server running on port 9002');
});
