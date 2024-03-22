"use client";

import { motion } from "framer-motion";

import AppNavbar from "~/components/app/navbar";

export default function AuthTemplate({
  logoSrc,
  logoBase64,
  children,
}: React.PropsWithChildren<{ logoSrc: string; logoBase64: string }>) {
  return (
    <motion.div
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      transition={{ duration: 0.5 }}
    >
      <AppNavbar logoSrc={logoSrc} logoBase64={logoBase64} />
      {children}
    </motion.div>
  );
}
