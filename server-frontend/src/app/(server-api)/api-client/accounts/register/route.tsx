import pool from "../../../../../lib/connect_model";

export async function POST(request: Request) {
    if (request.body === null) {
        return new Response(JSON.stringify({ error: "Request body is empty" }), { status: 400 });
    }

    const body = await request.json();
    const { username, email, password } = body;

    let error = "";
    const conn = await pool.getConnection();
    if (!conn) {
        return new Response(JSON.stringify({ error: "Database connection failed" }), { status: 500 });
    }

    try {
        const res = await conn.query('INSERT INTO users (username, email, password) VALUES (?, ?, ?)', [username, email, password]);
        console.log("Query result:", res);
    } catch (err) {
        console.error("Error executing query:", err);
        error = "User registration failed";
    } finally {
        if (conn) conn.release();
    }

    if (error) {
        return new Response(JSON.stringify({ error }), { status: 409 });
    }

    return new Response(JSON.stringify({ message: "User registered successfully" }), { status: 201 });
}