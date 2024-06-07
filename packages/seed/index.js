import mysql from "mysql2/promise"

const connection = await mysql.createConnection({
    host: "localhost",
    user: "user",
    password: "password",
    database: "restaurantFlow",
    port: 3306,
})

// A simple SELECT query
try {
    const [results, fields] = await connection.query(/*sql*/ `select * from dummyTable`)

    console.log(results) // results contains rows returned by server
    console.log(fields) // fields contains extra meta data about results, if available
} catch (err) {
    console.log(err)
}

connection.destroy()
