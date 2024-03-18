import { redirect } from "next/navigation";

import RegisterForm from "~/components/app/main/register";

import getSession from "~/server/auth";

export default function RegisterPage() {
  const session = getSession();

  if (session) redirect("/");

  return (
    <main className="flex min-h-screen flex-col items-center justify-center">
      <RegisterForm />
    </main>
  );
}
