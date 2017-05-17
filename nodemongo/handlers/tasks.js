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
        // store.getAll()
        //     .then(tasks => {
        //         res.json(tasks);
        //     })
        //     .catch(next);
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

    router.patch('/v1/tasks/:taskID', async (req, res, next) => {
        let taskID = req.params.taskID;
        try {
            let updatedTask = await store.setComplete(taskID, req.body.complete);
            res.json(updatedTask);
        } catch(err) {
            next(err);
        }
    });

    router.delete('/v1/tasks/:taskID', async (req, res, next) => {
        let taskID = req.params.taskID;
        try {
            await store.delete(taskID);
            res.send(`deleted task ${taskID}`);
        } catch(err) {
            next(err);
        }
    });

    return router;
};
