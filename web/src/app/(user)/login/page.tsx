import { redirect } from "next/navigation";

import LoginForm from "~/components/app/main/login";

import getAuth from "~/server/auth";

export default function LoginPage() {
  const session = getAuth();

  if (session) {
    redirect("/");
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-center">
      <LoginForm />
    </main>
  );
}
