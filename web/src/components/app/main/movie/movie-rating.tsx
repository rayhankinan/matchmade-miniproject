"use client";

import { useQuery, useQueryClient, useMutation } from "@tanstack/react-query";
import { Rating } from "@smastrom/react-rating";
import { ExclamationTriangleIcon } from "@radix-ui/react-icons";

import { Alert, AlertDescription, AlertTitle } from "~/components/ui/alert";
import { useToast } from "~/components/ui/use-toast";

import api from "~/client/api";

interface RatingResponse {
  data: number;
}

export default function MovieRating({
  id,
  enabled,
}: {
  id: number;
  enabled: boolean;
}) {
  const { toast } = useToast();

  const queryClient = useQueryClient();

  const getRatingQuery = useQuery({
    queryKey: ["movie-rating", id],
    queryFn: async () => await api.get<RatingResponse>(`/movies/rate/${id}`),
    enabled,
  });

  const {
    data: getRatingData,
    error: getRatingError,
    status: getRatingStatus,
  } = getRatingQuery;

  const changeRatingMutation = useMutation({
    mutationFn: async (data: { rating: number }) =>
      api.patch<RatingResponse>(`/movies/rate/${id}`, data),
    onSuccess: async () => {
      await queryClient.invalidateQueries({
        queryKey: ["movie-rating", id],
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

  const { mutate: changeRatingMutate, isPending: isChangeRatingPending } =
    changeRatingMutation;

  const handleRatingChange = (value: number) =>
    changeRatingMutate({ rating: value });

  if (!enabled || getRatingStatus === "pending")
    return <Rating value={0} isDisabled className="w-1/2" />;

  if (getRatingStatus === "error")
    return (
      <Alert variant="destructive" className="w-1/2">
        <ExclamationTriangleIcon className="h-4 w-4" />
        <AlertTitle>Error</AlertTitle>
        <AlertDescription>{getRatingError.message}</AlertDescription>
      </Alert>
    );

  return (
    <Rating
      value={getRatingData.data.data}
      onChange={handleRatingChange}
      isDisabled={isChangeRatingPending}
      className="w-1/2"
    />
  );
}
