"use strict";

const express = require('express');
const morgan = require('morgan');

const host = process.env.HOST || '';
const port = process.env.PORT || '80';

const app = express();

app.use(morgan(process.env.LOGFORMAT || 'dev'));

app.get("/hello", (req, res) => {
    let user = JSON.parse(req.header('X-User'));
    res.send(`Hello ${user.firstName} ${user.lastName}! (responded from port ${port})`);
});

app.listen(port, host, () => {
    console.log(`hello service is listening at http://${host}:${port}...`);
});
