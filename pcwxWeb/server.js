const express = require('express');
const app = express();
const path = require('path');

// Serve static files from the "public" directory
app.use(express.static(path.join(__dirname, 'public')));

app.listen(3002, () => {
    console.log('Server is running on http://8.130.44.136:3002');
});
