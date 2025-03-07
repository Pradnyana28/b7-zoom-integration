import { JSX } from 'react';

interface ButtonProps {
  children: string | JSX.Element;
  onClick?: () => void;
}

export default function Button({ children }: ButtonProps) {
  return (
    <button className="bg-blue-400 shadow-sm rounded-sm border-blue-800 p-2 text-white">
      {children}
    </button>
  );
}
