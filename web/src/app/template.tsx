"use client";

import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { motion } from "framer-motion";

import AppNavbar from "~/components/app/navbar";

export default function RootTemplate({ children }: React.PropsWithChildren) {
  return (
    <motion.div
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      transition={{ duration: 0.5 }}
    >
      <AppNavbar />
      {children}
      {process.env.NODE_ENV !== "production" && (
        <footer className="hidden md:block">
          <ReactQueryDevtools initialIsOpen={false} />
        </footer>
      )}
    </motion.div>
  );
}
