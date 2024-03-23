"use client";

import { useQuery, useQueryClient, useMutation } from "@tanstack/react-query";

import { Button } from "~/components/ui/button";
import { useToast } from "~/components/ui/use-toast";

import Spinner from "~/components/app/icon/spinner";
import api from "~/client/api";

interface IsMovieInWatchlistResponse {
  data: boolean;
}

export default function AddToWatchlistButton({
  id,
  title,
  posterPath,
  enabled,
}: {
  id: number;
  title: string;
  posterPath: string | null;
  enabled: boolean;
}) {
  const { toast } = useToast();

  const queryClient = useQueryClient();

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

  const addToWatchlistMutation = useMutation({
    mutationFn: async () =>
      api.post(`/movies/add`, {
        movieID: id,
        title,
        image: posterPath,
      }),
    onSuccess: async () => {
      // TODO: Invalidate search-watchlist query

      await queryClient.invalidateQueries({
        queryKey: ["is-movie-in-watchlist", id],
      });
    },
    onError: (error) => {
      toast({
        variant: "destructive",
        title: "An error occurred",
        description: error.message,
      });
    },
  });

  const { mutate: mutateAddToWatclist, isPending: isAddToWatchlistPending } =
    addToWatchlistMutation;

  const removeFromWatchlistMutation = useMutation({
    mutationFn: async () => api.delete(`/movies/remove/${id}`),
    onSuccess: async () => {
      // TODO: Invalidate watchlist query

      await queryClient.invalidateQueries({
        queryKey: ["is-movie-in-watchlist", id],
      });
    },
    onError: (error) => {
      toast({
        variant: "destructive",
        title: "An error occurred",
        description: error.message,
      });
    },
  });

  const {
    mutate: mutateRemoveFromWatchlist,
    isPending: isRemoveFromWatchlistPending,
  } = removeFromWatchlistMutation;

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
    <Button
      variant={isMovieInWatchlistData.data.data ? "destructive" : "default"}
      onClick={
        isMovieInWatchlistData.data.data
          ? () => mutateRemoveFromWatchlist()
          : () => mutateAddToWatclist()
      }
      disabled={
        isMovieInWatchlistData.data.data
          ? isRemoveFromWatchlistPending
          : isAddToWatchlistPending
      }
    >
      {isMovieInWatchlistData.data.data
        ? "Remove from watchlist"
        : "Add to watchlist"}
    </Button>
  );
}
