"use strict";

const express = require('express');
const Task = require('../models/tasks/task.js');

//export a function from this module 
//that accepts a tasks store implementation
module.exports = function(store) {
    //create a new Mux
    let router = express.Router();

    router.get('/v1/tasks', async (req, res, next) => {
        try {
            let tasks = await store.getAll();
            res.json(tasks);
        } catch(err) {
            next(err);
        }
    });

    router.post('/v1/tasks', async (req, res, next) => {
        try {
            let task = new Task(req.body);
            let err = task.validate();
            if (err) {
                res.status(400).send(err.message);
            } else {
                let result = await store.insert(task);
                res.json(task);
            }
        } catch(err) {
            next(err);
        }
    });

    return router;
};
