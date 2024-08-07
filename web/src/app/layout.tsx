import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { Inter } from "next/font/google";

import { Toaster } from "~/components/ui/toaster";

import ClientProvider from "~/providers/client";
import ThemeProvider from "~/providers/theme";
import SessionProvider from "~/providers/auth";
import getSession from "~/server/auth";
import { cn } from "~/lib/utils";

import "@smastrom/react-rating/style.css";
import "~/styles/globals.css";

const inter = Inter({
  subsets: ["latin"],
  variable: "--font-sans",
});

export const metadata = {
  title: "Movie Watchlist",
  description: "A simple movie watchlist app",
  icons: [{ rel: "icon", url: "/favicon.ico" }],
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const session = getSession();

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
            <SessionProvider value={session}>
              {children}
              {process.env.NODE_ENV !== "production" && (
                <footer className="hidden md:block">
                  <ReactQueryDevtools initialIsOpen={false} />
                </footer>
              )}
            </SessionProvider>
            <Toaster />
          </ClientProvider>
        </ThemeProvider>
      </body>
    </html>
  );
}
