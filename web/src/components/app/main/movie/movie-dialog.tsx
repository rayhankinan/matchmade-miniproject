"use client";

import { useCallback, useEffect, useState } from "react";
import { useSearchParams, usePathname, useRouter } from "next/navigation";

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "~/components/ui/dialog";

export default function MovieDialog({ id }: { id: string | undefined }) {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const router = useRouter();

  const [isOpen, setIsOpen] = useState(false);

  const onOpenChange = useCallback(
    (open: boolean) => {
      // Remove the jbv query param when the dialog is closed
      if (!open) {
        const params = new URLSearchParams(searchParams);
        params.delete("jbv");
        router.push(`${pathname}?${params.toString()}`);
      }

      setIsOpen(open);
    },
    [pathname, router, searchParams],
  );

  useEffect(() => {
    if (id !== undefined) setIsOpen(true);
  }, [id]);

  return (
    <Dialog open={isOpen} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>This is the title</DialogTitle>
          <DialogDescription>This is the description</DialogDescription>
        </DialogHeader>
        <DialogFooter>This is the footer</DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
