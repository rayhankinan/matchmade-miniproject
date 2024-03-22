"use client";

import { motion } from "framer-motion";

import AppNavbar from "~/components/app/navbar";
import SearchBar from "~/components/app/main/home/search-bar";

export default function MovieTemplate({
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
      <AppNavbar
        logoSrc={logoSrc}
        logoBase64={logoBase64}
        searchBar={<SearchBar />}
        displayProfile
      />
      {children}
    </motion.div>
  );
}
