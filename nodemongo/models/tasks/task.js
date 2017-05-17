"use strict";

class Task {
    constructor(props) {
        Object.assign(this, props);
    }

    validate() {
        if (!this.title) {
            return new Error('you must supply a title');
        }
    }
}

module.exports = Task;
