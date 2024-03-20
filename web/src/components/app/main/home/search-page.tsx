"use client";

import { Fragment, useCallback, useEffect, useRef } from "react";
import { useInfiniteQuery } from "@tanstack/react-query";
import { ExclamationTriangleIcon } from "@radix-ui/react-icons";

import { Alert, AlertDescription, AlertTitle } from "~/components/ui/alert";

import MoviePreview from "~/components/app/main/movie/movie-preview";
import Spinner from "~/components/app/icon/spinner";
import movieApi from "~/client/movie-api";

interface SearchMovieResponse {
  page: number;
  total_pages: number;
  total_results: number;
  results: {
    id: number;
    title: string;
    poster_path: string | null;
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

  if (status === "pending") return <Spinner />;

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
                data={{
                  id: movie.id,
                  title: movie.title,
                  posterPath: movie.poster_path,
                }}
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
