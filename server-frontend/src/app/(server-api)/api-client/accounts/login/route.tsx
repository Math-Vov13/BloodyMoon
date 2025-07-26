import pool from "../../../../../lib/connect_model";

export async function POST(request: Request) {
    if (request.body === null) {
        return new Response(JSON.stringify({ error: "Request body is empty" }), { status: 400 });
    }

    const body = await request.json();
    const { email, password } = body;

    let error = "";
    let res = [];
    const conn = await pool.getConnection();
    if (!conn) {
        return new Response(JSON.stringify({ error: "Database connection failed" }), { status: 500 });
    }

    try {
        res = await conn.query('SELECT * FROM users WHERE email = ? AND password = ?', [email, password]);
        console.log("Query result:", res);
    } catch (err) {
        console.error("Error executing query:", err);
        error = "User login failed";
    } finally {
        if (conn) conn.release();
    }

    if (error) {
        return new Response(JSON.stringify({ error }), { status: 404 });
    }

    if (res.length === 0) {
        return new Response(JSON.stringify({ error: "Invalid email or password" }), { status: 401 });
    }

    return new Response(JSON.stringify({ message: "User logged in successfully" }), { status: 200 });
}