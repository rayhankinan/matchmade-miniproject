"use client";

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
    </motion.div>
  );
}
