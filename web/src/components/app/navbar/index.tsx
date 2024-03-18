"use client";

import { Button } from "~/components/ui/button";

import ModeToggle from "~/components/app/navbar/mode-toggle";

export default function AppNavbar() {
  return (
    <header className="flex h-20 w-full shrink-0 items-center px-4 md:px-6">
      <div className="ml-auto flex gap-2">
        <Button variant="outline">Sign in</Button>
        <Button>Sign Up</Button>
        <ModeToggle />
      </div>
    </header>
  );
}
