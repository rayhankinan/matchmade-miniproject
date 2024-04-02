"use client";

import { useQuery, useQueryClient, useMutation } from "@tanstack/react-query";

import { Button } from "~/components/ui/button";
import { useToast } from "~/components/ui/use-toast";

import MovieRating from "~/components/app/main/movie/movie-rating";
import MovieTags from "~/components/app/main/movie/movie-tags";
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
  onChange,
}: {
  id: number;
  title: string;
  posterPath: string | null;
  enabled: boolean;
  onChange?: () => void;
}) {
  const { toast } = useToast();

  const queryClient = useQueryClient();

  const isMovieInWatchlistQuery = useQuery({
    queryKey: ["is-movie-in-watchlist", id],
    queryFn: async () =>
      await api.get<IsMovieInWatchlistResponse>(
        `/movies/watchlist/exist/${id}`,
      ),
    enabled,
  });

  const {
    data: isMovieInWatchlistData,
    error: isMovieInWatchlistError,
    status: isMovieInWatchlistStatus,
  } = isMovieInWatchlistQuery;

  const addToWatchlistMutation = useMutation({
    mutationFn: async (data: {
      movieID: number;
      title: string;
      image: string | null;
    }) => api.post(`/movies/add`, data),
    onSuccess: async () => {
      // Call onChange to update the UI
      if (onChange) onChange();

      // Invalidate watchlist query to improve UX
      await queryClient.invalidateQueries({
        queryKey: ["discover-watchlist"],
      });

      await queryClient.invalidateQueries({
        queryKey: ["search-watchlist"],
      });

      // Invalidate is-movie-in-watchlist query to update the button state
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

  const { mutate: mutateAddToWatchlist, isPending: isAddToWatchlistPending } =
    addToWatchlistMutation;

  const removeFromWatchlistMutation = useMutation({
    mutationFn: async () => await api.delete(`/movies/remove/${id}`),
    onSuccess: async () => {
      // Call onChange to update the UI
      if (onChange) onChange();

      // Invalidate watchlist query to improve UX
      await queryClient.invalidateQueries({
        queryKey: ["discover-watchlist"],
      });

      await queryClient.invalidateQueries({
        queryKey: ["search-watchlist"],
      });

      // Invalidate is-movie-in-watchlist query to update the button state
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

  if (isMovieInWatchlistStatus === "error")
    return (
      <Button disabled variant="secondary">
        Error occured: {isMovieInWatchlistError.message}
      </Button>
    );

  return (
    <>
      <MovieRating id={id} enabled={isMovieInWatchlistData.data.data} />
      <MovieTags id={id} enabled={isMovieInWatchlistData.data.data} />
      <Button
        variant={isMovieInWatchlistData.data.data ? "destructive" : "default"}
        onClick={
          isMovieInWatchlistData.data.data
            ? () => mutateRemoveFromWatchlist()
            : () =>
                mutateAddToWatchlist({
                  movieID: id,
                  title,
                  image: posterPath,
                })
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
    </>
  );
}
