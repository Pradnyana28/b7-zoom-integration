import './global.css';

export const metadata = {
  title: 'Zoom Integration',
  description: 'Create, update, view, and delete meetings',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
