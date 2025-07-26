import mariadb from "mariadb";

const pool = mariadb.createPool({
    port: parseInt(process.env.MARIA_DB_PORT || "3306", 10),
    host: process.env.MARIA_DB_HOST,
    user: process.env.MARIA_DB_USER,
    password: process.env.MARIA_DB_PASSWORD,
    database: process.env.MARIA_DB_DATABASE,
    connectionLimit: 5
});

export default pool;