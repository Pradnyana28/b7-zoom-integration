export default async function Index() {
  const userResponse = await fetch('http://localhost:3000/api/v1/users/me');
  const user = await userResponse.json();

  return (
    <div className="w-100 m-auto text-center flex flex-col items-center gap-6 p-4">
      <div className="w-full flex flex-row justify-between">
        <p className="text-xl font-bold">Zoom Account</p>
        <p className="text-xl font-bold">{user.first_name}!</p>
      </div>
    </div>
  );
}
