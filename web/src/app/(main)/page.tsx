import DiscoverMovie from "~/components/app/main/home/discover-page";
import SearchMovie from "~/components/app/main/home/search-page";
import MovieDialog from "~/components/app/main/movie/movie-dialog";

export default function HomePage({
  searchParams,
}: {
  searchParams: {
    q?: string;
    jbv?: string;
  };
}) {
  const { q: query, jbv: id } = searchParams;

  return (
    <main className="flex min-h-screen flex-col items-center justify-center gap-12">
      {query ? <SearchMovie query={query} /> : <DiscoverMovie />}
      <MovieDialog id={id} />
    </main>
  );
}
