import { RocketIcon } from "lucide-react";

import { Alert, AlertDescription, AlertTitle } from "~/components/ui/alert";

export default function WatchlistAlert() {
  return (
    <Alert className="w-1/3">
      <RocketIcon className="h-4 w-4" />
      <AlertTitle>Heads up!</AlertTitle>
      <AlertDescription>
        You need to log in to access your watchlist.
      </AlertDescription>
    </Alert>
  );
}
