"use client";

import { useQuery } from "@tanstack/react-query";

import { Button } from "~/components/ui/button";

import Spinner from "~/components/app/icon/spinner";
import api from "~/client/api";

interface IsMovieInWatchlistResponse {
  data: boolean;
}

export default function AddToWatchlistButton({
  id,
  enabled,
}: {
  id: number;
  enabled: boolean;
}) {
  const isMovieInWatchlistQuery = useQuery({
    queryKey: ["is-movie-in-watchlist", id],
    queryFn: async () =>
      api.get<IsMovieInWatchlistResponse>(`/movies/watchlist/exist/${id}`),
    enabled,
  });

  const {
    data: isMovieInWatchlistData,
    error: isMovieInWatchlistError,
    status: isMovieInWatchlistStatus,
  } = isMovieInWatchlistQuery;

  if (isMovieInWatchlistStatus === "pending")
    return (
      <Button disabled variant="secondary">
        <Spinner /> Loading
      </Button>
    );

  if (isMovieInWatchlistError)
    return (
      <Button disabled variant="secondary">
        Error occured: {isMovieInWatchlistError.message}
      </Button>
    );

  return (
    <Button variant={isMovieInWatchlistData.data ? "destructive" : "default"}>
      {isMovieInWatchlistData.data
        ? "Remove from watchlist"
        : "Add to watchlist"}
    </Button>
  );
}
