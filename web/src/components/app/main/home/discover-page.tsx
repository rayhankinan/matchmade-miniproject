"use client";

import { Fragment } from "react";
import { useInfiniteQuery } from "@tanstack/react-query";

import { Button } from "~/components/ui/button";

import movieApi from "~/client/movie-api";

interface DiscoverMovieResponse {
  page: number;
  results: {
    id: number;
    title: string;
    poster_path: string;
  }[];
}

export default function DiscoverMovie() {
  const discoverMovieQuery = useInfiniteQuery({
    queryKey: ["trending-movie"],
    queryFn: async ({ pageParam }) =>
      movieApi.get<DiscoverMovieResponse>(`/discover/movie?page=${pageParam}`),
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
  } = discoverMovieQuery;

  if (status === "pending") return <div>Loading...</div>;

  if (status === "error") return <div>Error: {error.message}</div>;

  return (
    <div className="container flex flex-col items-center justify-center gap-12 px-4 py-16">
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
      <Button
        onClick={() => fetchNextPage()}
        disabled={!hasNextPage || isFetchingNextPage}
      >
        {isFetchingNextPage
          ? "Loading more..."
          : hasNextPage
            ? "Load More"
            : "Nothing more to load"}
      </Button>
      <div>{isFetching && !isFetchingNextPage ? "Fetching..." : null}</div>
    </div>
  );
}
