"use strict";

const express = require('express');
const morgan = require('morgan');
const cors = require('cors');
const bodyParser = require('body-parser');
const mongodb = require('mongodb');

const TaskStore = require('./models/tasks/mongostore.js');

const port = process.env.PORT || 80;
const host = process.env.HOST || '';
const mongoAddr = process.env.MONGOADDR || 'localhost:27017';

//create an Experss application
const app = express();
//add request logging
app.use(morgan(process.env.LOGFORMAT || 'dev'));
//add CORS headers
app.use(cors());
//add middleware that parses
//any JSON posted to this app.
//the parsed data will be available
//on the req.body property
app.use(bodyParser.json());

mongodb.MongoClient.connect(`mongodb://${mongoAddr}/demo`)
    .then(db => {
        let colTasks = db.collection('tasks');
        let store = new TaskStore(colTasks);
        let handlers = require('./handlers/tasks.js');
        app.use(handlers(store));

        //error handler
        app.use((err, req, res, next) => {
            console.error(err);
            res.status(err.status || 500).send(err.message);
        });

        app.listen(port, host, () => {
            console.log(`server is listening at http://${host}:${port}...`);
        });
    })
    .catch(err => {
        console.error(err);
    });
