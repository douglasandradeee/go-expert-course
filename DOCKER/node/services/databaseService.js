const connection = require('../config/db');

const insertPerson = (name, callback) => {
    const sql = `INSERT INTO people(name) VALUES(?)`;
    connection.query(sql, [name], (err, results) => {
        if (err) {
            return callback(err, null);
        };
        callback(null, results);
    });
};

module.exports = {
    insertPerson
};