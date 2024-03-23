"use client";

import { useState } from "react";
import Link from "next/link";
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
import AddToWatchlistButton from "~/components/app/main/movie/add-to-watchlist-button";
import useSession from "~/hooks/auth";
import movieApi from "~/client/movie-api";

interface MovieDetailResponse {
  id: number;
  title: string;
  overview: string;
  release_date: string;
  poster_path: string | null;
  videos: {
    results: {
      site: "YouTube" | "Vimeo";
      key: string;
    }[];
  };
}

export default function MovieDialog({
  id,
  refreshOnChange = false,
  children,
}: React.PropsWithChildren<{ id: number; refreshOnChange?: boolean }>) {
  const [open, setOpen] = useState(false);

  const session = useSession();

  const detailQuery = useQuery({
    queryKey: ["movie-detail", id],
    queryFn: async () =>
      movieApi.get<MovieDetailResponse>(
        `/movie/${id}?append_to_response=videos`,
      ),
    enabled: open,
  });

  const {
    data: detailData,
    error: detailError,
    status: detailStatus,
  } = detailQuery;

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className="max-h-screen overflow-y-scroll">
        {detailStatus === "pending" ? (
          <div className="flex items-center justify-center">
            <Spinner />
          </div>
        ) : detailStatus === "error" ? (
          <>
            <DialogHeader>
              <DialogTitle>Error</DialogTitle>
            </DialogHeader>
            <Alert variant="destructive" className="w-1/2">
              <ExclamationTriangleIcon className="h-4 w-4" />
              <AlertTitle>Error</AlertTitle>
              <AlertDescription>{detailError.message}</AlertDescription>
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
                <div className="text-xl">{detailData.data.title}</div>
              </DialogTitle>
              <DialogDescription>
                <div className="text-md">
                  &quot;{detailData.data.overview}&quot;
                </div>
              </DialogDescription>
              <div className="flex items-center pt-2">
                <CalendarIcon className="mr-2 h-4 w-4 opacity-70" />{" "}
                <span className="text-xs text-muted-foreground">
                  Released on {detailData.data.release_date}
                </span>
              </div>
            </DialogHeader>
            {detailData.data.videos.results.length > 0 && (
              <div className="flex flex-col items-center justify-center gap-4">
                <Carousel className="w-full max-w-md">
                  <CarouselContent>
                    {detailData.data.videos.results.map((video) => (
                      <CarouselItem
                        key={video.key}
                        className="flex items-center justify-center"
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
            <DialogFooter className="md:justify-center">
              {session ? (
                <div className="flex flex-col items-center justify-center gap-4">
                  <AddToWatchlistButton
                    id={id}
                    title={detailData.data.title}
                    posterPath={detailData.data.poster_path}
                    enabled={open}
                    onChange={
                      refreshOnChange ? () => setOpen(false) : undefined
                    }
                  />
                </div>
              ) : (
                <Button variant="link">
                  <Link href="/login">Login to add to watchlist</Link>
                </Button>
              )}
            </DialogFooter>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
}
