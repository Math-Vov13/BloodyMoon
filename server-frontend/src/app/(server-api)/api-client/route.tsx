export async function GET() {
    // use Prisma client ?

    return new Response('Client API route is working', {
        status: 200,
        headers: {
            'Content-Type': 'text/plain',
        },
    });
}