export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
        <main>
            {children}
        </main>
        <footer className="text-center text-sm text-gray-500 mt-4">
            <p>© 2023 Your Company. All rights reserved.</p>
        </footer>
    </>
  );
}
