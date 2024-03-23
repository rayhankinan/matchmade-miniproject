import { RocketIcon } from "@radix-ui/react-icons";

import { Alert, AlertDescription, AlertTitle } from "~/components/ui/alert";

export default function LoginAlert() {
  return (
    <Alert className="w-1/2">
      <RocketIcon className="h-4 w-4" />
      <AlertTitle>Heads up!</AlertTitle>
      <AlertDescription>
        You are currently logged in. Please log out to access this page.
      </AlertDescription>
    </Alert>
  );
}
