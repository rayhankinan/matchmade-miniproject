"use client";

import { useCallback, useEffect, useState } from "react";
import { useSearchParams, usePathname, useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { ExclamationTriangleIcon } from "@radix-ui/react-icons";
import { CalendarIcon, LoaderIcon } from "lucide-react";

import { Alert, AlertDescription, AlertTitle } from "~/components/ui/alert";
import { Button } from "~/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "~/components/ui/dialog";

import movieApi from "~/client/movie-api";

interface MovieDetailResponse {
  id: number;
  title: string;
  overview: string;
  release_date: string;
  videos: {
    site: "YouTube" | "Vimeo";
    key: string;
  }[];
}

export default function MovieDialog({ id }: { id: string | undefined }) {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const router = useRouter();

  const detailQuery = useQuery({
    queryKey: ["movie-detail", id],
    queryFn: async () =>
      movieApi.get<MovieDetailResponse>(
        `/movie/${id}?append_to_response=videos`,
      ),
    enabled: id !== undefined,
  });

  const { data, error, status } = detailQuery;

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
        {status === "pending" ? (
          <div className="flex items-center justify-center">
            <LoaderIcon className="h-8 w-8 animate-spin" />
          </div>
        ) : status === "error" ? (
          <>
            <DialogHeader>
              <DialogTitle>Error</DialogTitle>
            </DialogHeader>
            <Alert variant="destructive">
              <ExclamationTriangleIcon className="h-4 w-4" />
              <AlertTitle>Error</AlertTitle>
              <AlertDescription>{error.message}</AlertDescription>
            </Alert>
            <DialogFooter className="sm:justify-start">
              <DialogClose asChild>
                <Button type="button" variant="secondary">
                  Close
                </Button>
              </DialogClose>
            </DialogFooter>
          </>
        ) : (
          <>
            <DialogHeader>
              <DialogTitle>{data.data.title}</DialogTitle>
              <DialogDescription>{data.data.overview}</DialogDescription>
              <div className="flex items-center pt-2">
                <CalendarIcon className="mr-2 h-4 w-4 opacity-70" />{" "}
                <span className="text-xs text-muted-foreground">
                  {data.data.release_date}
                </span>
              </div>
            </DialogHeader>
            <DialogFooter className="sm:justify-start">
              <DialogClose asChild>
                <Button type="button" variant="secondary">
                  Close
                </Button>
              </DialogClose>
            </DialogFooter>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
}
