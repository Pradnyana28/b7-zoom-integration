import AccountInfo from './_includes/components/accountInfo';
import Button from './_includes/components/button';

export default async function Index() {
  const userResponse = await fetch('http://localhost:3000/api/v1/users/me');
  const user = await userResponse.json();

  return (
    <div className="w-100 m-auto text-center flex flex-col items-center gap-6 p-4">
      <div className="w-full flex flex-row items-center justify-between">
        <AccountInfo user={user} />
        <Button>Add Meeting</Button>
      </div>

      <div className="w-full flex flex-row items-center justify-between bg-gray-100 p-4 rounded-md shadow-sm">
        <p>You don&apos;t have any meetings yet.</p>
      </div>
    </div>
  );
}
