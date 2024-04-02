"use client";

import { useQuery } from "@tanstack/react-query";
import { ExclamationTriangleIcon } from "@radix-ui/react-icons";

import { Alert, AlertDescription, AlertTitle } from "~/components/ui/alert";
import { Badge } from "~/components/ui/badge";

import api from "~/client/api";

interface TagsResponse {
  data: string[];
}

export default function MovieTags({
  id,
  enabled,
}: {
  id: number;
  enabled: boolean;
}) {
  const getTagsQuery = useQuery({
    queryKey: ["movie-tags", id],
    queryFn: async () => await api.get<TagsResponse>(`/movies/tags/${id}`),
    enabled,
  });

  const {
    data: getTagsData,
    error: getTagsError,
    status: getTagsStatus,
  } = getTagsQuery;

  if (!enabled || getTagsStatus === "pending") return null;

  if (getTagsStatus === "error")
    return (
      <Alert variant="destructive" className="w-1/2">
        <ExclamationTriangleIcon className="h-4 w-4" />
        <AlertTitle>Error</AlertTitle>
        <AlertDescription>{getTagsError.message}</AlertDescription>
      </Alert>
    );

  return (
    <div className="flex flex-row gap-4">
      {getTagsData.data.data.length > 0 &&
        getTagsData.data.data.map((tag) => <Badge key={tag}>{tag}</Badge>)}
    </div>
  );
}
