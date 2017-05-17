"use strict";

const express = require('express');
const morgan = require('morgan');
const cors = require('cors');
const bodyParser = require('body-parser');
const mongodb = require('mongodb');

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

//TODO: connect to the Mongo database
//add the tasks handlers
//and start listening for HTTP requests

