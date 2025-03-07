export async function GET() {
  try {
    const result = await fetch('http://localhost:8080/v1/users/me');
    const data = await result.json();

    return new Response(JSON.stringify(data), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    });
  } catch (error) {
    return new Response(
      JSON.stringify({
        message: (error as Error).message,
      }),
      {
        status: 421,
        headers: {
          'Content-Type': 'application/json',
        },
      }
    );
  }
}
