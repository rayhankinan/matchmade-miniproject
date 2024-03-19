"use client";

import Image from "next/image";
import { CalendarIcon } from "@radix-ui/react-icons";

import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "~/components/ui/hover-card";

import { env } from "~/env";

interface MoviePreviewProps {
  id: number;
  title: string;
  posterPath: string | null;
  overview: string;
  releaseDate: string;
}

export default function MoviePreview({
  title,
  posterPath,
  overview,
  releaseDate,
}: MoviePreviewProps) {
  return (
    <HoverCard>
      <HoverCardTrigger asChild>
        {posterPath ? (
          <Image
            src={`${env.NEXT_PUBLIC_MOVIE_IMAGE_URL}/w500${posterPath}`}
            blurDataURL={`${env.NEXT_PUBLIC_MOVIE_IMAGE_URL}/w92${posterPath}`}
            placeholder="blur"
            alt={title}
            width={200}
            height={300}
            className="h-auto w-auto cursor-pointer rounded-sm"
          />
        ) : (
          <div className="h-[375px] w-[250px] cursor-pointer rounded-sm bg-gray-300 bg-[url('/images/placeholder.svg')] bg-center bg-no-repeat"></div>
        )}
      </HoverCardTrigger>
      <HoverCardContent className="w-80">
        <div className="space-y-1">
          <h4 className="text-sm font-semibold">{title}</h4>
          <p className="text-sm">{overview}</p>
          <div className="flex items-center pt-2">
            <CalendarIcon className="mr-2 h-4 w-4 opacity-70" />{" "}
            <span className="text-xs text-muted-foreground">{releaseDate}</span>
          </div>
        </div>
      </HoverCardContent>
    </HoverCard>
  );
}
