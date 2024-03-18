import { Inter } from "next/font/google";

import ClientProvider from "~/providers/client";
import ThemeProvider from "~/providers/theme";
import AuthProvider from "~/providers/auth";
import { cn } from "~/lib/utils";

import "~/styles/globals.css";
import { getAuth } from "~/server/auth";

const inter = Inter({
  subsets: ["latin"],
  variable: "--font-sans",
});

export const metadata = {
  title: "Create T3 App",
  description: "Generated by create-t3-app",
  icons: [{ rel: "icon", url: "/favicon.ico" }],
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const session = getAuth();

  return (
    <html lang="en" suppressHydrationWarning>
      <body
        className={cn(
          "min-h-screen bg-background font-sans antialiased",
          inter.variable,
        )}
      >
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <ClientProvider>
            <AuthProvider session={session}>{children}</AuthProvider>
          </ClientProvider>
        </ThemeProvider>
      </body>
    </html>
  );
}
