import { LoaderIcon } from "lucide-react";

export default function Spinner() {
  return (
    <div className="p-4">
      <LoaderIcon className="h-8 w-8 animate-spin" />
    </div>
  );
}
