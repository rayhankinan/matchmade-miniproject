import Image from "next/image";

import MovieDialog from "~/components/app/main/movie/movie-dialog";
import { env } from "~/env";

export default function MovieClickable({
  data: { id, title, posterPath },
}: {
  data: {
    id: number;
    title: string;
    posterPath: string | null;
  };
}) {
  return (
    <MovieDialog id={id}>
      {posterPath ? (
        <div className="relative h-[375px] w-[250px] cursor-pointer transition-all hover:scale-[1.05] hover:shadow-md">
          <Image
            src={`${env.NEXT_PUBLIC_MOVIE_IMAGE_URL}/w500${posterPath}`}
            alt={title}
            fill
            sizes="(min-width: 500px) 50vw, 100vw"
            className="rounded-sm"
            priority
          />
        </div>
      ) : (
        <div className="h-[375px] w-[250px] cursor-pointer rounded-sm bg-gray-300 bg-[url('/images/placeholder.svg')] bg-center bg-no-repeat transition-all hover:scale-[1.05] hover:shadow-md"></div>
      )}
    </MovieDialog>
  );
}
