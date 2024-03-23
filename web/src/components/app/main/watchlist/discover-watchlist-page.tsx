"use client";

import { Fragment, useCallback, useEffect, useRef } from "react";
import { useInfiniteQuery } from "@tanstack/react-query";
import { ExclamationTriangleIcon } from "@radix-ui/react-icons";

import { Alert, AlertDescription, AlertTitle } from "~/components/ui/alert";

import MovieClickable from "~/components/app/main/movie/movie-clickable";
import Spinner from "~/components/app/icon/spinner";
import api from "~/client/api";

interface DiscoverWatchlistResponse {
  page: number;
  total_pages: number;
  total_results: number;
  data: {
    MovieID: number;
    Title: string;
    Image: {
      String: string;
      Valid: boolean;
    };
    Rating: {
      Int64: number;
      Valid: boolean;
    };
  }[];
}

export default function DiscoverWatchlist() {
  const discoverWatchlistQuery = useInfiniteQuery({
    queryKey: ["discover-watchlist"],
    queryFn: async ({ pageParam }) =>
      await api.get<DiscoverWatchlistResponse>(
        `/movies/watchlist?page=${pageParam}&pageSize=20`,
      ),
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
  } = discoverWatchlistQuery;

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
  }, [fetchData]);

  if (status === "pending") return <Spinner />;

  if (status === "error")
    return (
      <Alert variant="destructive" className="w-1/2">
        <ExclamationTriangleIcon className="h-4 w-4" />
        <AlertTitle>Error</AlertTitle>
        <AlertDescription>{error.message}</AlertDescription>
      </Alert>
    );

  return (
    <>
      <div className="container flex flex-row flex-wrap items-center justify-center gap-12">
        {data.pages.map((page, index) => (
          <Fragment key={index}>
            {page.data.data.map((movie) => (
              <MovieClickable
                key={movie.MovieID}
                data={{
                  id: movie.MovieID,
                  title: movie.Title,
                  posterPath: movie.Image.Valid ? movie.Image.String : null,
                }}
                refreshOnChange
              />
            ))}
          </Fragment>
        ))}
      </div>
      {isFetchingNextPage && <Spinner />}
      <div ref={observerTarget}></div>
    </>
  );
}
