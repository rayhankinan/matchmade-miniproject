"use client";

import { useSearchParams, usePathname, useRouter } from "next/navigation";
import Image from "next/image";
import { CalendarIcon } from "@radix-ui/react-icons";

import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "~/components/ui/hover-card";

import { env } from "~/env";

export default function MoviePreview({
  id,
  title,
  posterPath,
  overview,
  releaseDate,
}: {
  id: number;
  title: string;
  posterPath: string | null;
  overview: string;
  releaseDate: string;
}) {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const router = useRouter();

  const onClick = () => {
    const params = new URLSearchParams(searchParams);
    params.set("jbv", id.toString());
    router.push(`${pathname}?${params.toString()}`);
  };

  return (
    <HoverCard>
      <HoverCardTrigger asChild>
        {posterPath ? (
          <div
            className="relative h-[375px] w-[250px] cursor-pointer overflow-hidden"
            onClick={() => onClick()}
          >
            <Image
              src={`${env.NEXT_PUBLIC_MOVIE_IMAGE_URL}/w500${posterPath}`}
              alt={title}
              fill
              sizes="(min-width: 500px) 50vw, 100vw"
              className="rounded-sm object-cover"
              priority
            />
          </div>
        ) : (
          <div className="h-[375px] w-[250px] cursor-pointer rounded-sm bg-gray-300 bg-[url('/images/placeholder.svg')] bg-center bg-no-repeat"></div>
        )}
      </HoverCardTrigger>
      <HoverCardContent className="w-80">
        <div className="space-y-1">
          <h4 className="text-sm font-semibold">{title}</h4>
          <p className="text-sm">
            {overview.length <= 100 ? overview : `${overview.slice(0, 100)}...`}
          </p>
          <div className="flex items-center pt-2">
            <CalendarIcon className="mr-2 h-4 w-4 opacity-70" />{" "}
            <span className="text-xs text-muted-foreground">{releaseDate}</span>
          </div>
        </div>
      </HoverCardContent>
    </HoverCard>
  );
}
