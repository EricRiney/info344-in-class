"use strict";

const express = require('express');
const morgan = require('morgan');
const cors = require('cors');
const bodyParser = require('body-parser');
const mongodb = require('mongodb');

const port = process.env.PORT || 80;
const host = process.env.HOST || '';
const mongoAddr = process.env.MONGOADDR || 'localhost:27017';

const app = express();
app.use(morgan(process.env.LOGFORMAT || 'dev'));
app.use(cors());
app.use(bodyParser.json());

