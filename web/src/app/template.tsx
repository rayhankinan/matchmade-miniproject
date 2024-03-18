"use client";

import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

export default function RootTemplate({ children }: React.PropsWithChildren) {
  return (
    <div>
      {children}
      {process.env.NODE_ENV !== "production" && (
        <footer className="hidden md:block">
          <ReactQueryDevtools initialIsOpen={false} />
        </footer>
      )}
    </div>
  );
}
