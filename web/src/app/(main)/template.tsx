"use client";

import { motion } from "framer-motion";

import AppNavbar from "~/components/app/navbar";
import SearchBar from "~/components/app/main/home/search-bar";

export default function RootTemplate({ children }: React.PropsWithChildren) {
  return (
    <motion.div
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      transition={{ duration: 0.5 }}
    >
      <AppNavbar searchBar={<SearchBar />} displayProfile />
      {children}
    </motion.div>
  );
}
