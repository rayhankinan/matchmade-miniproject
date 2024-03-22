"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { ExclamationTriangleIcon } from "@radix-ui/react-icons";
import { CalendarIcon } from "lucide-react";

import { Alert, AlertDescription, AlertTitle } from "~/components/ui/alert";
import { Button } from "~/components/ui/button";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "~/components/ui/carousel";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "~/components/ui/dialog";

import Spinner from "~/components/app/icon/spinner";
import MovieTrailer from "~/components/app/main/movie/movie-trailer";
import movieApi from "~/client/movie-api";

interface MovieDetailResponse {
  id: number;
  title: string;
  overview: string;
  release_date: string;
  videos: {
    results: {
      site: "YouTube" | "Vimeo";
      key: string;
    }[];
  };
}

export default function MovieDialog({
  id,
  children,
}: React.PropsWithChildren<{ id: number }>) {
  const [open, setOpen] = useState(false);

  const detailQuery = useQuery({
    queryKey: ["movie-detail", id],
    queryFn: async () =>
      movieApi.get<MovieDetailResponse>(
        `/movie/${id}?append_to_response=videos`,
      ),
    enabled: open,
  });

  const { data, error, status } = detailQuery;

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent>
        {status === "pending" ? (
          <div className="flex items-center justify-center">
            <Spinner />
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
              <DialogTitle>
                <div className="text-xl">{data.data.title}</div>
              </DialogTitle>
              <DialogDescription>
                <div className="text-md">&quot;{data.data.overview}&quot;</div>
              </DialogDescription>
              <div className="flex items-center pt-2">
                <CalendarIcon className="mr-2 h-4 w-4 opacity-70" />{" "}
                <span className="text-xs text-muted-foreground">
                  Released on {data.data.release_date}
                </span>
              </div>
            </DialogHeader>
            {data.data.videos.results.length > 0 && (
              <div className="flex flex-col items-center justify-center">
                <Carousel className="w-full max-w-md">
                  <CarouselContent>
                    {data.data.videos.results.map((video) => (
                      <CarouselItem
                        key={video.key}
                        className="flex flex-col items-center justify-center"
                      >
                        <MovieTrailer
                          data={{
                            site: video.site,
                            key: video.key,
                          }}
                        />
                      </CarouselItem>
                    ))}
                  </CarouselContent>
                  <CarouselPrevious />
                  <CarouselNext />
                </Carousel>
              </div>
            )}
            <DialogFooter className="sm:justify-center">
              <DialogClose asChild>
                <Button type="button" variant="secondary">
                  Add to Watchlist
                </Button>
              </DialogClose>
            </DialogFooter>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
}
