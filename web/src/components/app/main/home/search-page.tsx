"use client";

import { Fragment, useCallback, useEffect, useRef } from "react";
import { ExclamationTriangleIcon } from "@radix-ui/react-icons";
import { useInfiniteQuery } from "@tanstack/react-query";

import { Alert, AlertDescription, AlertTitle } from "~/components/ui/alert";
import { LoadingSpinner } from "~/components/svg/spinner";

import movieApi from "~/client/movie-api";

interface SearchMovieResponse {
  page: number;
  results: {
    id: number;
    title: string;
    poster_path: string | null;
    release_date: string;
  }[];
}

export default function SearchMovie({ query }: { query: string }) {
  const searchMovieQuery = useInfiniteQuery({
    queryKey: ["search-movie", query],
    queryFn: async ({ pageParam }) =>
      movieApi.get<SearchMovieResponse>(
        `/search/movie?query=${query}&page=${pageParam}`,
      ),
    initialPageParam: 1,
    getNextPageParam: (currentResponse) => currentResponse.data.page + 1,
  });
  const {
    data,
    error,
    fetchNextPage,
    hasNextPage,
    isFetching,
    isFetchingNextPage,
    status,
  } = searchMovieQuery;

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

  return (
    <div className="container flex flex-col items-center justify-center gap-12 px-4 py-16">
      {status === "pending" ? (
        <LoadingSpinner />
      ) : status === "error" ? (
        <Alert variant="destructive">
          <ExclamationTriangleIcon className="h-4 w-4" />
          <AlertTitle>Error</AlertTitle>
          <AlertDescription>{error.message}</AlertDescription>
        </Alert>
      ) : (
        <>
          {data.pages.map((page, i) => (
            <Fragment key={i}>
              {page.data.results.map((movie) => (
                <div key={movie.id}>
                  <p>ID: {movie.id}</p>
                  <p>Title: {movie.title}</p>
                  <p>Poster Path: {movie.poster_path}</p>
                </div>
              ))}
            </Fragment>
          ))}
          {isFetchingNextPage && <LoadingSpinner />}
          <div ref={observerTarget}></div>
        </>
      )}
    </div>
  );
}
