"use strict";

const mongodb = require('mongodb'); //for mongodb.ObjectID()

/**
 * MongoStore is a concrete store for Task models
 */
class MongoStore {
    /**
     * Constructs a new MongoStore
     * @param {mongodb.Collection} collection 
     */
    constructor(collection) {
        this.collection = collection;
    }

    /**
     * getAll returns all tasks in the store
     */
    getAll() {
        //TODO: implement this
    }

    /**
     * insert inserts a new Task into the store
     * @param {Task} task 
     */
    insert(task) {
        //TODO: implement this
    }

    /**
     * setComplete sets the complete status of the task
     * @param {string} id 
     * @param {bool} complete 
     */
    async setComplete(id, complete) {
        //TODO: implement this
    }

    /**
     * delete deletes the task with the given object ID
     * @param {string} id 
     */
    delete(id) {
        //TODO: implement this
    }
}

//export the class
module.exports = MongoStore;
