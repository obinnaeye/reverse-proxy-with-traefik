const express = require('express');
const app = express();
const port = 8082;

app.get('/', (req, res) => {
    res.send('Hello, JS World!');
});

app.listen(port, () => {
    console.log(`Service2 is running on http://localhost:${port}`);
});
