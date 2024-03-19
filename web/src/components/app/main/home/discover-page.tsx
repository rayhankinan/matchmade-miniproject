"use client";

import { Fragment, useCallback, useEffect, useRef } from "react";
import { ExclamationTriangleIcon } from "@radix-ui/react-icons";
import { LoaderIcon } from "lucide-react";
import { useInfiniteQuery } from "@tanstack/react-query";

import { Alert, AlertDescription, AlertTitle } from "~/components/ui/alert";

import MoviePreview from "~/components/app/main/movie/movie-preview";
import movieApi from "~/client/movie-api";

interface DiscoverMovieResponse {
  page: number;
  total_pages: number;
  total_results: number;
  results: {
    id: number;
    title: string;
    poster_path: string | null;
    overview: string;
    release_date: string;
  }[];
}

export default function DiscoverMovie() {
  const discoverMovieQuery = useInfiniteQuery({
    queryKey: ["trending-movie"],
    queryFn: async ({ pageParam }) =>
      movieApi.get<DiscoverMovieResponse>(`/discover/movie?page=${pageParam}`),
    initialPageParam: 1,
    getNextPageParam: (currentResponse) =>
      currentResponse.data.page < currentResponse.data.total_pages
        ? currentResponse.data.page + 1
        : undefined,
  });

  const {
    data,
    error,
    fetchNextPage,
    hasNextPage,
    isFetching,
    isFetchingNextPage,
    status,
  } = discoverMovieQuery;

  const observerTarget = useRef<HTMLDivElement>(null);

  const fetchData = useCallback(async () => {
    if (isFetching || !hasNextPage) return;

    await fetchNextPage();
  }, [isFetching, hasNextPage, fetchNextPage]);

  useEffect(() => {
    const currentTarget = observerTarget.current;

    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0] && entries[0].isIntersecting) void fetchData();
      },
      { threshold: 1 },
    );

    if (currentTarget) observer.observe(currentTarget);

    return () => {
      if (currentTarget) observer.unobserve(currentTarget);
    };
  }, [observerTarget, fetchData]);

  if (status === "pending")
    return <LoaderIcon className="h-8 w-8 animate-spin" />;

  if (status === "error")
    return (
      <Alert variant="destructive">
        <ExclamationTriangleIcon className="h-4 w-4" />
        <AlertTitle>Error</AlertTitle>
        <AlertDescription>{error.message}</AlertDescription>
      </Alert>
    );

  return (
    <>
      <div className="container flex flex-row flex-wrap items-center justify-center gap-12">
        {data.pages.map((page, i) => (
          <Fragment key={i}>
            {page.data.results.map((movie) => (
              <MoviePreview
                key={movie.id}
                id={movie.id}
                title={movie.title}
                posterPath={movie.poster_path}
                overview={movie.overview}
                releaseDate={movie.release_date}
              />
            ))}
          </Fragment>
        ))}
      </div>
      {isFetchingNextPage && <LoaderIcon className="h-8 w-8 animate-spin" />}
      <div ref={observerTarget}></div>
    </>
  );
}
