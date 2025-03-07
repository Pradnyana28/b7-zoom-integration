interface AccountInfoProps {
  user: {
    first_name: string;
    last_name: string;
    email: string;
  };
}

export default function AccountInfo({ user }: AccountInfoProps) {
  return (
    <div className="w-full flex flex-col justify-between items-start">
      <p className="text-xl font-bold">Zoom Account</p>
      <p className="text-sm">
        {user.first_name} {user.last_name}
      </p>
    </div>
  );
}
